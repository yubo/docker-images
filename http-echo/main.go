package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		log.Printf("%s %s", r.Method, r.URL.String())

		var buf bytes.Buffer
		if err := json.Indent(&buf, b, " >", "  "); err != nil {
			log.Printf("json decode err: %s", err)
			return
		}
		log.Println(buf.String())
	})))
}
