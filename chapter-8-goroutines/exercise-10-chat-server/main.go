// TODO improvements: broadcast incoming and leaving users

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type message struct {
	sender string
	text   string
}

type client struct {
	name  string
	inMsg chan *message
	conn  net.Conn
}
type clients map[string]*client

const (
	PORT                      = ":8000"
	DISCONNECT_CLIENT_TIMEOUT = time.Minute * 10
)

var mut = sync.Mutex{}

func main() {
	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("TCP server work on port %s\n", PORT)

	clientMap := make(clients)
	broadcast := make(chan *message)

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

func acceptConnect(conn net.Conn, clientMap clients, br chan<- *message) {
	sendGreetMessage(conn, clientMap)

	scanner := bufio.NewScanner(conn)
	clientName := conn.RemoteAddr().String()
	for scanner.Scan() {
		clientName = scanner.Text()
		mut.Lock()
		if _, ok := clientMap[clientName]; ok {
			fmt.Fprintf(conn, "The name %s has already taken. Try another name: ", clientName)
			mut.Unlock()
			continue
		} else {
			log.Printf("User \"%s\" connection success\n", clientName)
			fmt.Fprintf(conn, "Hi, %s!\n", clientName)
			mut.Unlock()
			break
		}
	}

	msg := make(chan *message)
	defer close(msg)

	mut.Lock()
	clientMap[clientName] = &client{conn: conn, inMsg: msg, name: clientName}
	mut.Unlock()

	defer func() {
		mut.Lock()
		delete(clientMap, clientName)
		mut.Unlock()
	}()

	go runWriter(clientMap[clientName])

	timer := time.NewTimer(DISCONNECT_CLIENT_TIMEOUT)
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Fprintln(conn, "Timout disconnect...")
		log.Printf("User \"%s\" timeout disconnect\n", clientName)
		conn.Close()
	}()

	for scanner.Scan() {
		timer.Reset(DISCONNECT_CLIENT_TIMEOUT)
		br <- &message{sender: clientName, text: scanner.Text()}
	}

	log.Printf("User \"%s\" disconnected\n", clientName)
}

func runWriter(c *client) {
	buf := make([]*message, 0, 10)
	retryAttempts := 0
	retryTimer := time.NewTimer(0)
	retryTimer.Stop()

	defer retryTimer.Stop()

	sendBufferedMessages := func() {
		for idx, msg := range buf {
			if err := printMessage(c, msg); err == nil {
				buf = buf[idx+1:]
				retryAttempts = 0
			} else {
				if retryAttempts >= 10 {
					c.conn.Close()
					return
				}
				retryTimer.Reset(time.Second)
				retryAttempts++
				break
			}
		}
	}

	for {
		select {
		case m, ok := <-c.inMsg:
			if !ok {
				return
			}
			if len(buf) >= cap(buf) {
				c.conn.Close()
				return
			}
			retryTimer.Stop()
			buf = append(buf, m)
			sendBufferedMessages()
		case _, ok := <-retryTimer.C:
			if ok {
				sendBufferedMessages()
			}
		}
	}
}

func runBroadcasting(clientMap clients, br <-chan *message) {
	for msg := range br {
		mut.Lock()
		for _, c := range clientMap {
			c.inMsg <- msg
		}
		mut.Unlock()
	}
}

func sendGreetMessage(conn net.Conn, c clients) {
	mut.Lock()
	defer mut.Unlock()

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

func printMessage(c *client, msg *message) error {
	if c.name == msg.sender {
		return nil
	}
	_, err := fmt.Fprintf(c.conn, "[%s] %s\n", msg.sender, msg.text)

	return err
}
