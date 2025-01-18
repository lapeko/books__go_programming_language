package main

import (
	"bytes"
	"fmt"
	"github.com/lapeko/books__go_programming_language/chapter-7-interfaces/num"
	"io"
	"log"
)

type limitReader struct {
	originReader io.Reader
	limit        int64
	sentBytes    int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if lr.sentBytes >= lr.limit {
		return 0, io.EOF
	}
	chunkSize := num.Min(lr.limit-lr.sentBytes, int64(len(p)))
	n, err = lr.originReader.Read(p[:chunkSize])
	lr.sentBytes += int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{originReader: r, limit: n}
}

func main() {
	reader := bytes.NewBufferString("There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text.")
	lr := LimitReader(reader, 100)
	buf := make([]byte, 50)

	for {
		_, err := lr.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err.Error())
			}
			break
		}
		fmt.Print(string(buf))
	}
	fmt.Println()
}
