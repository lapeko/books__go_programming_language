package testutils

import (
	"reflect"
	"testing"
)

func EqualType(t *testing.T, got, expected any) {
	t.Helper()
	type1 := reflect.TypeOf(got)
	type2 := reflect.TypeOf(expected)
	if type1 != type2 {
		t.Errorf("Expected type %v, got %v", type1, type2)
	}
}
