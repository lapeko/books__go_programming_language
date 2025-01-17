package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	var counter int
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error during scanning: %v", err)
	}

	*c += WordCounter(counter)

	return counter, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	var counter int

	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error during scanning: %v", err)
	}

	*c += LineCounter(counter)

	return counter, nil
}

func main() {
	//var c ByteCounter
	//c.Write([]byte("Hi there"))
	//fmt.Println(c)
	//c.Write([]byte("123"))
	//fmt.Println(c)
	//c = 0
	//c.Write([]byte("123"))
	//fmt.Println(c)

	var w WordCounter
	fmt.Fprint(&w, "qasasdasdw asd asdasd we qweqwe")
	fmt.Println(w)

	var l LineCounter
	fmt.Fprint(&l, "qasasdasdw asd\nasdasd we qweqwe")
	fmt.Println(l)
}
