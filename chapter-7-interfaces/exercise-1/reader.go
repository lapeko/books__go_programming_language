package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

// [DONE] My own reader
// Ридер для симуляции больших данных
// Объединение Reader-ов
// Перевод текста в верхний регистр

type myReader struct {
	src []byte
}

func newMyReader(s []byte) *myReader {
	return &myReader{
		src: s,
	}
}

func (mr *myReader) Read(s []byte) (n int, err error) {
	if len(s) == 0 {
		return 0, errors.New("provided buffer is empty")
	}
	if len(mr.src) == 0 {
		return 0, io.EOF
	}

	readySize := copy(s, mr.src)
	mr.src = mr.src[readySize:]
	return readySize, nil
}

func main() {
	mr := newMyReader([]byte("Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of \"de Finibus Bonorum et Malorum\" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, \"Lorem ipsum dolor sit amet..\", comes from a line in section 1.10.32."))
	buf := make([]byte, 4)
	for i := 0; i < 5; i++ {
		_, err := mr.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err.Error())
			}
			break
		}
		fmt.Println(string(buf))
	}
}
