package main

import (
	"fmt"
	"github.com/lapeko/books__go_programming_language/chapter-10-testing/polindrome/utils"
)

func main() {
	fmt.Printf("%t\n", utils.IsPalindrome("АЛЛА"))
	fmt.Printf("%t\n", utils.IsPalindrome("шалаш"))
}
