package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mut     sync.Mutex
	counter = 0
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/counter", handleCount)
	mux.HandleFunc("/", handleMain)
	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	mut.Lock()
	counter++
	mut.Unlock()
	_, _ = fmt.Fprintf(w, "Requested url is: %q\n", r.URL.Path)
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Query[%q]=%q\n", k, v)
	}
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}

func handleCount(w http.ResponseWriter, req *http.Request) {
	mut.Lock()
	count := counter
	mut.Unlock()
	_, _ = fmt.Fprintf(w, "counter: %d", count)
}
