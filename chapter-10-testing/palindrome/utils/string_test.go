package utils

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input  string
		expect bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},
		{"desserts", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.expect {
			t.Errorf(`IsPalindrome("%s") = %t`, test.input, got)
		}
	}
}
