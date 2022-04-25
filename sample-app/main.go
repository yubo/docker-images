package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync/atomic"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

func main() {
	flag.Parse()

	var httpReqs int64
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "Content-Type: text/plain; version=0.0.4; charset=utf-8")
		fmt.Fprintf(w, "# HELP http_requests_total How many HTTP requests processed\n")
		fmt.Fprintf(w, "# TYPE http_requests_total counter\n")
		fmt.Fprintf(w, "http_requests_total %d\n", atomic.AddInt64(&httpReqs, 1))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&httpReqs, 1)
	})
	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&httpReqs, 1)
		x := 0.00001
		for i := 0; i <= 10000000; i++ {
			x += math.Sqrt(x)
		}
		w.Write([]byte("OK"))
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
