package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 1; i <= 4; i++ {
		runGoroutine(i, wg)
	}

	wg.Wait()
}

func runGoroutine(num int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		fmt.Printf("goroutine %d started\n", num)
		time.Sleep(time.Millisecond * getRandomDuration())
		fmt.Printf("goroutine %d done\n", num)
		wg.Done()
	}()
}

func getRandomDuration() time.Duration {
	return time.Duration(rand.Intn(1000))
}
