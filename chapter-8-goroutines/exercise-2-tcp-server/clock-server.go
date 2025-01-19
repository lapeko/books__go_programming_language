package main

import (
	"log"
	"net"
	"time"
)

func main() {
	const port = ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("tcp server run on port %s\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			log.Println("close connection failure")
		}
	}(conn)
	for {
		now := time.Now()
		withMilliseconds := now.Truncate(time.Millisecond).Format("15:04:05.000\n")
		if _, err := conn.Write([]byte(withMilliseconds)); err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
