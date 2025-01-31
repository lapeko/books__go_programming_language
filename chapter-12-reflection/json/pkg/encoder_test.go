package pkg

import (
	"reflect"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	e := newEncoder()
	eT := reflect.TypeOf(e)
	expectT := reflect.TypeOf(&encoder{})
	if eT != expectT {
		t.Errorf("Encoder onstance has a wrong type %q when %q expected", eT, expectT)
	}
}

func TestEncodeArray(t *testing.T) {
	tests := []struct {
		testArray any
		expected  string
	}{
		{
			[...]string{"val1", "test"},
			`{ "val1", "test" }`,
		},
		{
			[...]int{1, 3, 2, 5},
			`{ 1, 3, 2, 5 }`,
		},
	}

	for _, tt := range tests {
		e := newEncoder()
		e.encodeIterable(reflect.ValueOf(tt.testArray))
		res := e.String()
		if res != tt.expected {
			t.Errorf("encodeIterable(%v) = %q when %q expected", tt.testArray, res, tt.expected)
		}
	}
}

func TestEncodeSlice(t *testing.T) {
	tests := []struct {
		testSlice any
		expected  string
	}{
		{
			[]string{"val1", "test"},
			`{ "val1", "test" }`,
		},
		{
			[]int{1, 3, 2, 5},
			`{ 1, 3, 2, 5 }`,
		},
	}

	for _, tt := range tests {
		e := newEncoder()
		e.encodeIterable(reflect.ValueOf(tt.testSlice))
		res := e.String()
		if res != tt.expected {
			t.Errorf("encodeIterable(%v) = %q when %q expected", tt.testSlice, res, tt.expected)
		}
	}
}

func TestEncodeMap(t *testing.T) {
	tests := []struct {
		testMap  any
		expected string
	}{
		{
			map[string]string{"key1": "test"},
			`{ "key1": "test" }`,
		},
		{
			map[struct{ field int }]string{{field: 11}: "test3"},
			`{ ( field: 11 ): "test3" }`,
		},
	}

	for _, tt := range tests {
		e := newEncoder()
		e.encodeMap(reflect.ValueOf(tt.testMap))
		res := e.String()
		if res != tt.expected {
			t.Errorf("encodeMap(%v) = %q when %q expected", tt.testMap, res, tt.expected)
		}
	}
}
