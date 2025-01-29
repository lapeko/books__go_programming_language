package utils

import (
	"reflect"
)

func EqualType(got, expected any) bool {
	type1, type2 := reflect.TypeOf(got), reflect.TypeOf(expected)
	return type1 == type2
}
