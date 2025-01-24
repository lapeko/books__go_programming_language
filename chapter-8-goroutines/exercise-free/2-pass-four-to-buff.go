package main

import (
	"fmt"
	"time"
)

func main() {
	const size = 4
	ch := make(chan int, 3)
	done := make(chan struct{})

	go func() {
		for i := 1; i <= size; i++ {
			ch <- i
			fmt.Println("Inserting value: ", i)
		}
	}()

	go func() {
		for i := 0; i < size; i++ {
			time.Sleep(time.Second)
			fmt.Println("Read value: ", <-ch)
		}
		close(done)
	}()

	<-done
}
