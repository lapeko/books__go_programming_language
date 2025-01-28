package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var i interface{} = 3
	fmt.Println(reflect.TypeOf(i).String())
	var reader io.Reader = os.Stdin
	fmt.Println(reflect.TypeOf(reader).String())
	fmt.Println(reflect.TypeOf(3).String())
}
