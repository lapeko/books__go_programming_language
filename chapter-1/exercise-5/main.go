package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFunc)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Requested url is: %q", r.URL.Path)
}
