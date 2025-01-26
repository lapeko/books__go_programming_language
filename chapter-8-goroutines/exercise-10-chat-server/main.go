package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	inMsg chan string
	conn  net.Conn
}
type clients map[string]*client

const (
	PORT                      = ":8000"
	DISCONNECT_CLIENT_TIMEOUT = time.Minute * 10
)

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
	sendGreetMessage(conn, clientMap)

	scanner := bufio.NewScanner(conn)
	var clientName string
	for scanner.Scan() {
		clientName = scanner.Text()
		if _, ok := clientMap[clientName]; ok {
			fmt.Fprintf(conn, "The name %s has already taken. Try another name: ", clientName)
			continue
		} else {
			log.Printf("User %s connection sucess\n", clientName)
			fmt.Fprintf(conn, "Hi, %s!\n", clientName)
			break
		}
	}

	msg := make(chan string)
	defer close(msg)
	defer delete(clientMap, clientName)

	c := &client{conn: conn, inMsg: msg}
	clientMap[clientName] = c
	go runWriter(c)

	timer := time.NewTimer(DISCONNECT_CLIENT_TIMEOUT)
	go func() {
		<-timer.C
		fmt.Fprintln(conn, "Timout disconnect...")
		log.Printf("User %s timeout disconnect\n", clientName)
		conn.Close()
	}()

	for scanner.Scan() {
		timer.Reset(DISCONNECT_CLIENT_TIMEOUT)
		br <- scanner.Text()
	}

	log.Printf("User %s disconnected\n", clientName)
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

func sendGreetMessage(conn net.Conn, c clients) {
	fmt.Fprint(conn, "Connected clients: [")
	if len(c) != 0 {
		fmt.Fprintln(conn)
	}
	for addr := range c {
		fmt.Fprintf(conn, "\t\"%s\"\n", addr)
	}
	fmt.Fprintln(conn, "]")
	fmt.Fprint(conn, "Enter your client name: ")
}
