package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()

		if errors.Is(err, io.EOF) {
			break
		}

		fmt.Print(string(r))
	}
}
