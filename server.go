package main

import (
	"log"
	"net"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {

	defer c.Close()
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		fmt.Print(command)
		/*switch command {
			case "search":
				fmt.Println("A vida é bela") 
		}*/
	}
}