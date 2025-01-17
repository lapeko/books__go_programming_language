package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func newTree(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	getSortedSlice(values[:0], root)
	return root
}

func getSortedSlice(values []int, t *tree) []int {
	if t != nil {
		values = getSortedSlice(values, t.left)
		values = append(values, t.value)
		values = getSortedSlice(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	return fmt.Sprintf("%v", getSortedSlice(nil, t))
}

func main() {
	tree := newTree([]int{4, 33, 1, 6, 7, 2, 12})
	fmt.Println(tree)
}
