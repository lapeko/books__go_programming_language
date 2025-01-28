package utils

import (
	"regexp"
	"strings"
)

func IsPalindrome(text string) bool {
	re := regexp.MustCompile(`[^\p{L}\p{N}]`)
	cleaned := re.ReplaceAllString(text, "")
	runes := []rune(strings.ToLower(cleaned))
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-i-1] {
			return false
		}
	}
	return true
}
