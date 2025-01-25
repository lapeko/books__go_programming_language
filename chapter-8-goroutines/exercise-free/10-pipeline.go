package main

import "fmt"

func main() {
	naturals := make(chan int)
	shutdown := make(chan struct{})
	squares := make(chan int)

	go genNaturals(naturals, shutdown)
	go squareNaturals(naturals, squares, shutdown)
	runPipeTill(10_000, squares, shutdown)
}

func genNaturals(nts chan<- int, shn <-chan struct{}) {
	defer close(nts)
	for i := 1; ; i++ {
		select {
		case <-shn:
			return
		case nts <- i:
		}
	}
}

func squareNaturals(nts <-chan int, sqs chan<- int, shn <-chan struct{}) {
	defer close(sqs)
	for {
		select {
		case <-shn:
			return
		case n, ok := <-nts:
			if !ok {
				return
			}
			sqs <- n * n
		}
	}
}

func runPipeTill(maxNum int, sqs <-chan int, shn chan<- struct{}) {
	for {
		sq := <-sqs
		if sq > maxNum {
			close(shn)
			return
		}
		fmt.Println(sq)
	}
}
