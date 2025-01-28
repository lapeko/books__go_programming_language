package utils

func IsPalindrome(text string) bool {
	n := len(text)
	for i := 0; i < n/2; i++ {
		if text[i] != text[n-i-1] {
			return false
		}
	}
	return true
}
