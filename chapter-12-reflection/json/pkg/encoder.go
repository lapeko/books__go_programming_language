package pkg

import (
	"bytes"
	"fmt"
	"reflect"
)

type encoder struct {
	b *bytes.Buffer
}

func Marshal(data any) string {
	b := newEncoder()
	b.encode(reflect.ValueOf(data))
	return b.String()
}

func newEncoder() *encoder {
	return &encoder{&bytes.Buffer{}}
}

func (e *encoder) encode(v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		return
	case reflect.Bool:
		e.WriteString(fmt.Sprintf("%t", v.Bool()))
	case reflect.Uint, reflect.Int, reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
		e.WriteString(fmt.Sprintf("%d", v.Int()))
	case reflect.Float32, reflect.Float64:
		e.WriteString(fmt.Sprintf("%f", v.Float()))
	case reflect.String:
		fmt.Fprintf(e.b, "%q", v.String())
	case reflect.Struct:
		e.encodeStruct(v)
	case reflect.Array, reflect.Slice:
		e.encodeIterable(v)
	case reflect.Map:
		e.encodeMap(v)
	case reflect.Ptr:
		e.encodePtr(v)
	}
}

func (e *encoder) encodeStruct(v reflect.Value) {
	e.WriteString("{ ")
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if i != 0 {
			e.WriteString(", ")
		}
		f := t.Field(i)
		e.WriteString(fmt.Sprintf("%s: ", f.Name))
		e.encode(v.Field(i))
	}
	e.WriteString(" }")
}

func (e *encoder) encodeIterable(v reflect.Value) {
	e.WriteString("{ ")
	for i := 0; i < v.Len(); i++ {
		e.encode(v.Index(i))
		if i+1 < v.Len() {
			e.WriteString(", ")
		} else {
			e.WriteString(" ")
		}
	}
	e.WriteString("}")
}

func (e *encoder) encodeMap(v reflect.Value) {
	e.WriteString("{ ")
	keys := v.MapKeys()
	for i, key := range keys {
		e.encode(key)
		e.WriteString(`: `)
		e.encode(v.MapIndex(key))
		if i+1 < len(keys) {
			e.WriteString(", ")
		} else {
			e.WriteString(" ")
		}
	}
	e.WriteString("}")
}

func (e *encoder) encodePtr(v reflect.Value) {
	if v.IsNil() {
		e.WriteString("nil")
	}
	e.encode(v.Elem())
}

func (e *encoder) WriteString(s string) {
	e.b.WriteString(s)
}

func (e *encoder) String() string {
	return e.b.String()
}

type Actor struct {
	Role string
	Name string
}

type Movie struct {
	Title    string
	Subtitle string
	Year     int
	Actors   []Actor
	Oscars   []string
	Sequel   *string
}
