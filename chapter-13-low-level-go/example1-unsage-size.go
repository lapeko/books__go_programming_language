package main

import (
	"fmt"
	"unsafe"
)

const intSize = unsafe.Sizeof(0) // computed at compilation

type Good struct {
	a int64 // 8
	b int32 // 4
	c int8  // 2
} // Size 16

type Bad struct {
	a int8  // 2
	b int64 // 8
	c int32 // 4
} // Size 24

type Test struct {
	s   string
	num int8
}

func main() {
	var numPtr *int = nil
	// fmt.Println(*numPtr) // Panic
	fmt.Println(unsafe.Sizeof(*numPtr)) // Don't panic (expression not called)

	fmt.Println(unsafe.Sizeof(Good{})) // 16
	fmt.Println(unsafe.Sizeof(Bad{}))  // 24

	fmt.Println(unsafe.Offsetof(Good{}.c)) // starts at 12
	fmt.Println(unsafe.Offsetof(Bad{}.c))  // starts at 16
}
