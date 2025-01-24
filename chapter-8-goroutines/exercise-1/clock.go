package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleRes(conn)
	}
}

func handleRes(c net.Conn) {
	defer func() {
		fmt.Println("Connection closed")
		c.Close()
	}()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05:000\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
