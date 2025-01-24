package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer func() {
			time.Sleep(time.Second)
			close(done)
		}()

		counter := 0
		for v := range ch {
			fmt.Println(v)
			counter++
			if counter > 2 {
				close(ch)
			}
		}
	}()

	<-done
}
