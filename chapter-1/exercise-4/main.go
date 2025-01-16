package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = fmt.Sprintf("%s%s", "http://", url)
		}
		res, err := http.Get(url)
		if err != nil {
			log.Fatalln(fmt.Errorf("http get error: %w", err))
		}
		_, err = io.Copy(os.Stdout, res.Body)
		if err := res.Body.Close(); err != nil {
			log.Println(fmt.Errorf("http body close error: %w", err))
		}
		if err != nil {
			log.Fatalln(fmt.Errorf("http body read error: %w", err))
		}
	}
}
