package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	go animate(50)
	size := flag.Int("d", 10, "fibonacci size")
	flag.Parse()
	res := fibonacci(*size)
	fmt.Printf("\rfibonacci of %d is %d\n", *size, res)
}

func fibonacci(num int) int {
	if num < 2 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func animate(durationMS time.Duration) {
	for {
		for _, ch := range "-\\|/" {
			fmt.Printf("\r%c", ch)
			time.Sleep(time.Duration(time.Millisecond * durationMS))
		}
	}
}
