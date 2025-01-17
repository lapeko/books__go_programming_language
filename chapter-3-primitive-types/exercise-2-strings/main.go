package main

import "fmt"

func main() {
	s := "Здрасте"
	prefix := "Здра"

	for idx, char := range s {
		fmt.Printf("s[%d] %%c: %[2]c, %%v: %[2]v\n", idx, char)
	}

	fmt.Printf("%s hasPrefix %s: %t\n", s, prefix, hasPrefix(s, prefix))
}

func hasPrefix(s string, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	if s[:len(prefix)] != prefix {
		return false
	}
	return true
}
