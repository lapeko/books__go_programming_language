package main

import (
	"fmt"
	"github.com/lapeko/books__go_programming_language/chapter-2-program-structure/exercise-2-temperatures/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("20K is %vC\n", tempconv.KToC(tempconv.Kalvin(20)))
}
