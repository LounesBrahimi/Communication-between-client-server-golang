package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	continu := true
	for continu {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if string(message) != "A" {
			fmt.Println("Message recepted =====> ", message)
			conn.Write([]byte(string(message)))
		} else {
			conn.Close()
			continu = false
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ": 100")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server is listening on port 100")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}
		fmt.Println("Accept a connexion")
		go handleConnection(conn)
	}
}
