package main

import (
	"fmt"
	"github.com/lapeko/books__go_programming_language/intset/pkg/intset"
	map_engine "github.com/lapeko/books__go_programming_language/intset/pkg/intset/engine/map-engine"
)

func main() {
	set := intset.New(map_engine.New())
	fmt.Println(set.Has(100))
	set.Add(0)
	set.Add(20)
	set.Add(50)
	fmt.Println(set.String())
	fmt.Println(set.Has(50))
	set.Delete(50)
	fmt.Println(set.Has(50))
}
