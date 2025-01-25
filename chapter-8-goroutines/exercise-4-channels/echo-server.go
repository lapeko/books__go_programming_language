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

const duration = time.Second * 10

func handleConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	text := make(chan string)
	timer := time.NewTimer(duration)

	defer conn.Close()
	defer timer.Stop()
	defer fmt.Println("client disconnected")

	fmt.Println("new connection accepted")

	go func() {
		for scanner.Scan() {
			text <- scanner.Text()
		}
		close(text)
	}()

	for {
		select {
		case <-timer.C:
			close(text)
		case t, ok := <-text:
			if !ok {
				return
			}
			timer.Reset(duration)
			go echo(conn, t)
		}
	}
}

func echo(conn net.Conn, text string) {
	fmt.Println("text: ", text)
	fmt.Fprintln(conn, strings.ToUpper(text))
	time.Sleep(time.Second)
	fmt.Fprintln(conn, fmt.Sprintf("%s%s", strings.ToUpper(text[:1]), strings.ToLower(text[1:])))
	time.Sleep(time.Second)
	fmt.Fprintln(conn, strings.ToLower(text))
}
