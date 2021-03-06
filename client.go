package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:100")
	if err != nil {
		fmt.Println("Error of connexion")
		return
	}
	fmt.Println("Connected to localhost:100")
	reader := bufio.NewReader(os.Stdin)
	continu := true
	for continu {
		fmt.Println("Enter a line of text (empty to stop)")
		msg, _, _ := reader.ReadLine()
		if string(msg) == "A" {
			fmt.Fprintf(conn, fmt.Sprint(string(msg)))
			continu = false
			conn.Close()
			break
		} else {
			fmt.Println("Sending of :", string(msg))
			message := string(msg) + "\n"
			fmt.Fprintf(conn, fmt.Sprint(string(message)))
			reponse, _ := bufio.NewReader(conn).ReadString('\n')
			reponse = strings.TrimSuffix(reponse, "\n")
		}
	}
}
