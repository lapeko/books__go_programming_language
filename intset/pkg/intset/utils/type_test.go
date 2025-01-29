package utils

import (
	"os"
	"testing"
)

func TestEqualType(t *testing.T) {
	tests := []struct {
		a      any
		b      any
		expect bool
	}{
		{a: 1, b: 2, expect: true},
		{a: 1, b: os.Stdin, expect: false},
	}

	for _, tt := range tests {
		res := EqualType(tt.a, tt.b)
		if res != tt.expect {
			t.Errorf("EqualType(%T, %T) return %t", tt.a, tt.b, res)
		}
	}
}
