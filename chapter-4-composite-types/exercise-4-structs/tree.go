package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) add(value int) *tree {
	target := &t.right
	if value < t.value {
		target = &t.left
	}
	if *target == nil {
		*target = &tree{value: value}
	} else {
		(*target).add(value)
	}
	return t
}

func (t *tree) getValues() (values []int) {
	if t != nil {
		values = append(values, t.left.getValues()...)
		values = append(values, t.value)
		values = append(values, t.right.getValues()...)
	}

	return
}

func main() {
	fmt.Println((&tree{value: 50}).
		add(24).
		add(75).
		add(13).
		add(78).
		getValues())
}
