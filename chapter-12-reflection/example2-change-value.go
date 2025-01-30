package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 3
	v := reflect.ValueOf(&a)
	e := v.Elem()
	fmt.Println(e.CanAddr())
	ep := e.Addr().Interface().(*int)
	*ep = 10
	fmt.Println(a)
	e.SetInt(12)
	fmt.Println(a)
	e.Set(reflect.ValueOf(4))
	fmt.Println(a)
}
