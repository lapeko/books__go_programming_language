package main

import (
	"fmt"
	"math/big"
)

const (
	_ = 1 << (10 * iota)
	Kb
	Mb
	Gb
	Tb
	Pb
	Eb
)

var (
	Zb = new(big.Int).Mul(big.NewInt(Eb), big.NewInt(Kb))
	Yb = new(big.Int).Mul(Zb, big.NewInt(Kb))
)

func main() {
	fmt.Printf("Kb: %d\n", Kb)
	fmt.Printf("Mb: %d\n", Mb)
	fmt.Printf("Gb: %d\n", Gb)
	fmt.Printf("Tb: %d\n", Tb)
	fmt.Printf("Pb: %d\n", Pb)
	fmt.Printf("Eb: %d\n", Eb)
	fmt.Printf("Zb: %s\n", Zb)
	fmt.Printf("Yb: %s\n", Yb)
}
