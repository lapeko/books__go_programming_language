package main

import "fmt"

func main() {
	natural := make(chan int)
	squares := make(chan int)

	go genNaturals(natural)
	go squareNaturals(natural, squares)
	printSquares(squares)

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		natural <- i
	//	}
	//	close(natural)
	//}()
	//
	//go func() {
	//	for {
	//		nat, ok := <-natural
	//		if !ok {
	//			break
	//		}
	//		squares <- nat * nat
	//	}
	//	close(squares)
	//}()
	//
	//for square := range squares {
	//	fmt.Println(square)
	//}
}

func genNaturals(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squareNaturals(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printSquares(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
