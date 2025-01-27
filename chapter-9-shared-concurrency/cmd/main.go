package main

import (
	"fmt"
	"github.com/lapeko/books__go_programming_language/chapter-9-shared-concurrency/internal"
	"github.com/lapeko/books__go_programming_language/chapter-9-shared-concurrency/utils"
	"os"
	"time"
)

func main() {
	input, terminate := make(chan string), make(chan struct{})
	go utils.ScanStdin(os.Stdin, input, terminate)

	c := internal.NewCache(internal.BodyCacheFunc)
	go c.Serve()

	for url := range input {
		n := time.Now()
		res := c.Get(url, nil)
		fmt.Fprintf(os.Stderr, "Request to %s\nTime: %dms\n", url, time.Since(n).Milliseconds())
		if err := res.Err; err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		} else {
			fmt.Println(res.Body)
		}
	}
}
