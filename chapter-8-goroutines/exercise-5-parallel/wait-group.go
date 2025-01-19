package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numChan := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go heavyCalc(numChan, &wg)
	}

	go func() {
		wg.Wait()
		close(numChan)
	}()

	var sum int
	for num := range numChan {
		sum += num
	}
	fmt.Println("Sum is: ", sum)
}

func heavyCalc(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	random := int(rand.Int31n(1000))
	time.Sleep(time.Millisecond * time.Duration(random))
	out <- random
}
