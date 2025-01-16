package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			log.Fatalln(fmt.Errorf("http get error: %w", err))
		}
		body, err := io.ReadAll(res.Body)
		if err := res.Body.Close(); err != nil {
			log.Println(fmt.Errorf("http body close error: %w", err))
		}
		if err != nil {
			log.Fatalln(fmt.Errorf("http body read error: %w", err))
		}
		fmt.Printf("%s", body)
	}
}
