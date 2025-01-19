package main

import "fmt"

func main() {
	natural := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			natural <- i
		}
		close(natural)
	}()

	go func() {
		for {
			nat, ok := <-natural
			if !ok {
				break
			}
			squares <- nat * nat
		}
		close(squares)
	}()

	for square := range squares {
		fmt.Println(square)
	}
}
