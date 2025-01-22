package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		BuildPolygon(w)
	})
	log.Printf("Server is working on port %s\n", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
