package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(sqlQuote(nil))
	log.Println(sqlQuote(1))
	log.Println(sqlQuote(false))
	log.Println(sqlQuote("string here"))
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
