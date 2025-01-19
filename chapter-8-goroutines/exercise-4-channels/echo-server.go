package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	const port = ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("server is running on port %s\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept connection error. Skipping it...")
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new connection accepted")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		go echo(conn, scanner.Text())
	}

	fmt.Println("completed")
}

func echo(conn net.Conn, text string) {
	fmt.Println("text: ", text)
	fmt.Fprintln(conn, strings.ToUpper(text))
	time.Sleep(time.Second)
	fmt.Fprintln(conn, fmt.Sprintf("%s%s", strings.ToUpper(text[:1]), strings.ToLower(text[1:])))
	time.Sleep(time.Second)
	fmt.Fprintln(conn, strings.ToLower(text))
}
