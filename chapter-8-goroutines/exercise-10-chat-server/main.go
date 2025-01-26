package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	inMsg chan string
	conn  net.Conn
}
type clients map[string]*client

const PORT = ":8000"

func main() {
	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("TCP server work on port %s\n", PORT)

	clientMap := make(clients)
	broadcast := make(chan string)

	go runBroadcasting(clientMap, broadcast)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go acceptConnect(conn, clientMap, broadcast)
	}
}

func acceptConnect(conn net.Conn, clientMap clients, br chan<- string) {
	msg := make(chan string)
	defer close(msg)

	sendConnectedClients(conn, clientMap)

	addr := conn.RemoteAddr().String()

	if c, ok := clientMap[addr]; ok {
		c.conn.Close()
	}
	c := &client{conn: conn, inMsg: msg}
	clientMap[addr] = c
	go runWriter(c)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		br <- scanner.Text()
	}

	fmt.Println("Exit")
}

func runWriter(c *client) {
	for m := range c.inMsg {
		fmt.Fprintln(c.conn, m)
	}
}

func runBroadcasting(clientMap clients, br <-chan string) {
	for msg := range br {
		for _, c := range clientMap {
			c.inMsg <- msg
		}
	}
}

func sendConnectedClients(conn net.Conn, c clients) {
	fmt.Fprint(conn, "Connected clients: [")
	if len(c) != 0 {
		fmt.Fprintln(conn)
	}
	for addr := range c {
		fmt.Fprintf(conn, "\t%s\n", addr)
	}
	fmt.Fprintln(conn, "]")
}
