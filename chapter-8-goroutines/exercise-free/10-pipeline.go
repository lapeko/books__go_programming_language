package main

import "fmt"

func main() {
	naturals := make(chan int)
	shutdown := make(chan struct{})
	squares := make(chan int)

	go func() {
		defer close(naturals)
		for i := 1; ; i++ {
			select {
			case <-shutdown:
				return
			case naturals <- i:
			}
		}
	}()

	go func() {
		for n := range naturals {
			squares <- n * n
		}
	}()

	for {
		s := <-squares
		if s > 10_000 {
			close(shutdown)
			return
		}
		fmt.Println(s)
	}
}
