package main

import "fmt"

func main() {
	const size = 10
	ch := make(chan int)
	exit := make(chan struct{})

	go func() {
		for i := 0; i < size; i++ {
			ch <- i + 1
		}
	}()

	go func() {
		for i := 0; i < size; i++ {
			fmt.Println(<-ch)
		}
		close(exit)
	}()

	<-exit
}
