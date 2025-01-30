package main

import (
	"fmt"

	"github.com/lapeko/books__go_programming_language/intset/pkg/intset"
	binary_engine "github.com/lapeko/books__go_programming_language/intset/pkg/intset/binary-engine"
)

func main() {
	set := intset.New(binary_engine.New())
	fmt.Println(set.Has(100))
	set.Add(0)
	set.Add(20)
	set.Add(50)
	fmt.Println(set.String())
	fmt.Println(set.Has(50))
	set.Delete(50)
	fmt.Println(set.Has(50))
}
