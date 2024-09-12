package main

import (
	"log"
	"net"
	"bufio"
	"fmt"
	"strings"
)

func main() {

	//escuta na porta 8000 (pode ser monitorado com lsof -Pn -i4 | grep 8000)
	listener, err := net.Listen("tcp", "150.165.42.160:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		//aceita uma conexão criada por um cliente
		conn, err := listener.Accept()
		if err != nil {
			// falhas na conexão. p.ex abortamento
			log.Print(err)
			continue
		}
		// serve a conexão estabelecida
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {

	defer c.Close()
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		switch command {
			case "search":
				fmt.Println("A vida é bela") 
		}
	}
}