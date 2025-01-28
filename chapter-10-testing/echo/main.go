package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	out     io.Writer = os.Stdin
	newLine           = flag.Bool("n", false, "new line")
	divider           = flag.String("s", " ", "args divider symbol")
)

func main() {
	flag.Parse()
	echo(out, flag.Args(), *newLine, divider)
}

func echo(out io.Writer, args []string, newLine bool, divider *string) {
	fmt.Fprint(out, strings.Join(args, *divider))
	if newLine {
		fmt.Fprintln(out)
	}
}
