// Exercises 1.1 - 1.3
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	const interactions = 10_000 - 1
	start := time.Now()
	for i := 0; i < interactions; i++ {
		strings.Join(os.Args, ", ")
	}
	fmt.Println(strings.Join(os.Args, ", "))
	fmt.Println("JOIN took: ", time.Since(start).Nanoseconds())

	start = time.Now()
	var args string
	for i := 0; i < interactions; i++ {
		for key, value := range os.Args {
			args += "key: " + strconv.Itoa(key) + ", value: " + value + "\n"
		}
	}
	args = ""
	for key, value := range os.Args {
		args += "key: " + strconv.Itoa(key) + ", value: " + value + "\n"
	}
	fmt.Println(args)
	fmt.Println("For join took:", time.Since(start).Nanoseconds())
}
