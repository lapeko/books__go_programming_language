package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountWithTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountWithTableInLoop(x uint64) int {
	popCount := 0
	for i := 0; i < 8; i++ {
		popCount += int(pc[byte(x>>(i*8))])
	}
	return popCount
}

func PopCountInLoop(x uint64) int {
	popCount := 0
	for x > 0 {
		popCount += int(x & 1)
		x >>= 1
	}
	return popCount
}

func PopCountUsingBitClear(x uint64) int {
	popCount := 0
	for x > 0 {
		popCount++
		x &= x - 1
	}
	return popCount
}

func main() {
	randomNum := rand.Uint64()

	runTest(randomNum, PopCountWithTable)
	runTest(randomNum, PopCountWithTableInLoop)
	runTest(randomNum, PopCountInLoop)
	runTest(randomNum, PopCountUsingBitClear)
}

func runTest(randomNum uint64, fn func(x uint64) int) {
	start := time.Now()
	for i := 0; i < 100_000_000; i++ {
		fn(randomNum)
	}
	res := fn(randomNum)
	duration := time.Since(start)
	fmt.Printf("randNum: %d, number of non zero bytes: %d, duration: %d nS\n", randomNum, res, duration.Nanoseconds())
}
