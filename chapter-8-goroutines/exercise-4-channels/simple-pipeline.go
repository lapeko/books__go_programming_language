package main

import "fmt"

func main() {
	natural := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			natural <- i
		}
		close(natural)
	}()

	go func() {
		for num := range natural { // Читаем из канала `natural` до его закрытия
			squares <- num * num
		}
		close(squares)
	}()

	for square := range squares {
		fmt.Println(square)
	}
}
