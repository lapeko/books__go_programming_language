package main

import "fmt"

func String(a any) string {
	type stringer interface {
		String() string
	}
	switch a := a.(type) {
	case stringer:
		return a.String()
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", a)
	case float32, float64:
		return fmt.Sprintf("%f", a)
	case string:
		return a
	case bool:
		return fmt.Sprintf("%t", a)
	default:
		return "What?"
	}
}
