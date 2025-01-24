package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strings"
)

var cwd = pwd()

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("incomming request failed: %v", err)
			continue
		}
		handleRequest(conn)
	}
}

func handleRequest(c net.Conn) {
	defer func() {
		fmt.Println("closing connect...")
		c.Close()
	}()

	printPwd(c)

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		input := scanner.Text()
		inputWords := strings.Split(input, " ")
		cmd, arg := inputWords[0], inputWords[1:]
		switch cmd {
		case "ls":
			ls(arg, c)
		case "cd":
			cd(arg, c)
		case "get":
			get(arg, c)
		case "close":
			closeConnection(c)
		}
	}
}

func ls(args []string, c net.Conn) {
	defer printPwd(c)

	lsPath := cwd
	if len(args) != 0 {
		lsPath = path.Join(cwd, args[0])
	}

	cmd := exec.Command("ls", lsPath)

	output, err := cmd.Output()

	if os.IsNotExist(err) {
		c.Write([]byte(fmt.Sprintf("Directory %s does not exist\n", lsPath)))
		return
	} else if err != nil {
		c.Write([]byte(err.Error()))
		return
	}

	c.Write(output)
}

func cd(args []string, c net.Conn) {
	defer printPwd(c)

	if len(args) == 0 {
		c.Write([]byte("usage: cd <argument>\n"))
		return
	}

	newCwd := path.Join(cwd, args[0])

	if stat, err := os.Stat(newCwd); os.IsNotExist(err) {
		c.Write([]byte(fmt.Sprintf("Directory %s does not exist\n", newCwd)))
		return
	} else if !stat.IsDir() {
		c.Write([]byte(fmt.Sprintf("Path %s is not a folder\n", newCwd)))
		return
	} else if err != nil {
		c.Write([]byte(err.Error()))
		return
	}

	cwd = newCwd
}

func get(args []string, c net.Conn) {
	defer printPwd(c)

	if len(args) == 0 {
		c.Write([]byte("usage: get <argument>\n"))
		return
	}

	getFilePath := path.Join(cwd, args[0])
	stat, err := os.Stat(getFilePath)

	if os.IsNotExist(err) {
		c.Write([]byte(fmt.Sprintf("file %s not found\n", getFilePath)))
		return
	} else if stat.IsDir() {
		c.Write([]byte(fmt.Sprintf("path %s not a file\n", getFilePath)))
		return
	}

	cmd := exec.Command("cat", path.Join(cwd, args[0]))
	output, err := cmd.Output()

	if err != nil {
		c.Write([]byte(err.Error()))
		return
	}

	c.Write(output)
}

func closeConnection(c net.Conn) {
	c.Write([]byte("Goodbye!\n"))
	c.Close()
}

func pwd() string {
	dir, _ := os.Getwd()
	return dir
}

func printPwd(c net.Conn) {
	c.Write([]byte(fmt.Sprintf("%s: ", cwd)))
}
