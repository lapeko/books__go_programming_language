package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	slice := arr[1 : len(arr)-1]
	fmt.Printf("slice: %v, cap: %d, len: %d\n", slice, cap(slice), len(slice))
}
