package utils

import (
	"bufio"
	"io"
)

func ScanStdin(reader io.Reader, in chan<- string, terminate <-chan struct{}) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		select {
		case <-terminate:
			return
		default:
			in <- scanner.Text()
		}
	}
}
