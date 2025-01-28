package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		newLine   bool
		divider   string
		args      []string
		expResult string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
	}

	for _, test := range tests {
		buff := new(bytes.Buffer)
		echo(buff, test.args, test.newLine, &test.divider)
		if buff.String() != test.expResult {
			t.Errorf(`techo(buff, %v, %t, %s) printed %s when %s expected`, test.args, test.newLine, test.divider, buff.String(), test.expResult)
		}
	}
}
