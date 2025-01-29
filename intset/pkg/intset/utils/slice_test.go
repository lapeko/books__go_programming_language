package utils

import "testing"

func TestCompare(t *testing.T) {
	tests := []struct {
		name           string
		slice1, slice2 []int
		expect         bool
	}{
		{name: "Same slices", slice1: []int{1}, slice2: []int{1}, expect: true},
		{name: "Different slice lengths", slice1: make([]int, 1), slice2: make([]int, 2), expect: false},
		{name: "Different slices", slice1: []int{1}, slice2: []int{2}, expect: false},
	}

	for _, tt := range tests {
		res := Compare(tt.slice1, tt.slice2)
		if res != tt.expect {
			t.Errorf("%s error. Expect %t. Received %t", tt.name, tt.expect, res)
		}
	}
}
