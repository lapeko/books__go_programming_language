package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	const port = ":8000"
	conn, err := net.Dial("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	go func() {
		mustCopy(os.Stdout, conn)
		fmt.Println("Done")
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	fmt.Println("Complete")
}

func mustCopy(writer io.Writer, reader io.Reader) {
	if _, err := io.Copy(writer, reader); err != nil {
		log.Fatalln(err)
	}
}
