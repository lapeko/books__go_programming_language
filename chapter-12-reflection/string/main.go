package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var w io.Writer = os.Stdin
	v := reflect.ValueOf(w)
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(v.Kind())
	fmt.Println(v.String())
	fmt.Println(v.Type())
	fmt.Println(v.Interface())
	fmt.Println(v.Type() == reflect.TypeOf(os.Stdin))
}
