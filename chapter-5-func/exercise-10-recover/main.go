package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch r {
			case "expected":
				fmt.Println("expected error occurred")
			default:
				fmt.Fprintf(os.Stderr, "unexpected error occurred: %q\n", r)
			}
		}
	}()

	panicErrorMessage := flag.String("panic", "", "provide panic error message")
	flag.Parse()

	fmt.Printf("Provided panicError message: \"%s\"\n", *panicErrorMessage)

	if *panicErrorMessage != "" {
		panic(*panicErrorMessage)
	}

	fmt.Println("Main done")
}
