// Finds identical lines across multiple files and lists the files where each line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := map[string]map[string]bool{}
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("No files provided in args. Exit process...")
		return
	} else {
		readLinesFromFiles(args, counts)
	}

	for key, set := range counts {
		if len(set) > 1 {
			fmt.Print("Key: ", key, " is met in files: ")
			for key, _ := range set {
				fmt.Print(key, " ")
			}
			fmt.Println()
		}
	}
}

func readLinesFromFiles(files []string, counts map[string]map[string]bool) {
	for _, fileName := range files {
		file, err := os.Open(fileName)

		if err != nil {
			fmt.Println("Read file failure", err)
			continue
		}

		reader := bufio.NewScanner(file)
		for reader.Scan() {
			text := reader.Text()
			if counts[text] == nil {
				counts[text] = map[string]bool{}
			}
			counts[text][fileName] = true
		}
		if err := reader.Err(); err != nil {
			fmt.Println("Read file error:", err)
		}

		err = file.Close()
		if err != nil {
			fmt.Println("Close file failure", err)
			return
		}
	}
}
