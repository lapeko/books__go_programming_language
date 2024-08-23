package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := map[string]int{}
	args := os.Args[1:]

	if len(args) < 1 {
		countLines(os.Stdin, counts)
	} else {
		readLinesFromFiles(args, counts)
	}

	fmt.Printf("%v", counts)
}

func readLinesFromFiles(files []string, counts map[string]int) {
	for _, fileName := range files {
		file, err := os.Open(fileName)

		if err != nil {
			fmt.Println("Read file failure", err)
			continue
		}

		countLines(file, counts)

		err = file.Close()
		if err != nil {
			fmt.Println("Close file failure", err)
			return
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		counts[reader.Text()]++
	}
	if err := reader.Err(); err != nil {
		fmt.Println("Read file error:", err)
	}

}
