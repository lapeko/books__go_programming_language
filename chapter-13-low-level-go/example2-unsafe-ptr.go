package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var fl float64 = 1
	fmt.Printf("%#016x\n", *(*int)(unsafe.Pointer(&fl)))
}
