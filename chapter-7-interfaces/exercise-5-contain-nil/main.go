package main

import (
	"bytes"
	"io"
	"log"
	"reflect"
)

func main() {
	var buf *bytes.Buffer
	log.Printf("buf *bytes.Buffer == nil: %t\n", buf == nil)
	var w io.Writer = buf
	log.Printf("(w io.Writer = buf) == nil: %t\n", w == nil)

	dynamicType := reflect.TypeOf(w)
	dynamicValue := reflect.ValueOf(w)
	log.Printf("dynamic type of w: %s\n", dynamicType)
	log.Printf("dynamic value of w: %v\n", dynamicValue)
}
