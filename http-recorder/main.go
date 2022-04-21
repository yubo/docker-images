package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type recorder struct {
	config
	sync.Mutex
	reqs []*request
}

type request struct {
	*http.Request
	body []byte
}

func (p *request) String() string {
	b := &bytes.Buffer{}

	if !p.Close {
		p.body, _ = io.ReadAll(p.Body)
	}

	fmt.Fprintf(b, "%s \"%s %s %s\" %d \"%s\" \"%s\"\n%s",
		p.RemoteAddr,
		p.Method, p.URL.Path, p.Proto,
		len(p.body),
		p.Header.Get("Referer"),
		p.Header.Get("User-Agent"),
		string(p.body))
	return b.String()
}

func (p *recorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		p.dump(w, r)
		return
	}

	p.record(w, r)
}

func (p *recorder) dump(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	p.Lock()
	defer p.Unlock()

	for _, r := range p.reqs {
		fmt.Fprintf(w, "%s\n\n", r)
	}
}

func (p *recorder) record(w http.ResponseWriter, r *http.Request) {
	p.Lock()
	req := &request{Request: r}
	p.reqs = append(p.reqs, req)
	if len(p.reqs) > p.recordSize {
		a := p.reqs
		_, a = a[0], a[1:]
		p.reqs = a
	}
	p.Unlock()

	if p.debug {
		fmt.Printf("%s\n\n", req)
	}
}

func envDef(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}

	return def
}

type config struct {
	debug      bool
	recordSize int
	serverAddr string
}

func main() {
	var cf config
	flag.BoolVar(&cf.debug, "debug", false, "debug")
	flag.IntVar(&cf.recordSize, "size", 64, "record size")
	flag.StringVar(&cf.serverAddr, "addr", ":8080", "server address")

	flag.Parse()

	if os.Getenv("DEBUG") != "" {
		cf.debug = true
	}

	if addr := os.Getenv("SERVER_ADDRESS"); addr != "" {
		cf.serverAddr = addr
	}

	log.Fatal(http.ListenAndServe(cf.serverAddr, &recorder{config: cf}))
}
