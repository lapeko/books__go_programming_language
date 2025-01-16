package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	ch := make(chan int)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan int) {
	fmt.Printf("fetch start: %s\n", time.Now().Format("15:04:05.000"))
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("%s%s", "http://", url)
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(fmt.Errorf("http get error: %w", err))
	}
	ch <- res.StatusCode
	//_, err = io.Copy(os.Stdout, res.Body)
	//if err := res.Body.Close(); err != nil {
	//	log.Println(fmt.Errorf("http body close error: %w", err))
	//}
	fmt.Printf("fetch finish: %s\n", time.Now().Format("15:04:05.000"))
}
