package main

import (
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
	log.Printf("tcp client connectedn to port %s\n", port)
	mustCopy(os.Stdout, conn)
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatalln(err)
	}
}
