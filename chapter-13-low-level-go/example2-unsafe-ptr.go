package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var fl float64 = 1
	fmt.Printf("%#016x\n", *(*int)(unsafe.Pointer(&fl)))

	s := struct {
		a bool
		b int16
		c []int
	}{a: false, b: 16, c: []int{1, 3, 5}}

	pc := (*[]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.c)))
	*pc = append(*pc, 88)
	*pc = append(*pc, 99)
	fmt.Println(*pc)

	uptr := uintptr(unsafe.Pointer(&s.c[0]))
	for i := 0; i < len(s.c); i++ {
		fmt.Println(*(*int)(unsafe.Pointer(uptr)))
		uptr += unsafe.Sizeof(0)
	}
}
