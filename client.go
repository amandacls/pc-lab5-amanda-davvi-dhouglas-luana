package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./client search $file_hash")
		return
	}

	command := os.Args[1]
	fileHash := os.Args[2]

	if command != "search" {
		fmt.Println("Invalid command. Only 'search' is supported.")
		return
	}

	address := "localhost:8000"

	// Conecta ao servidor TCP
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Envia o comando ao servidor
	message := fmt.Sprintf("%s %s\n", command, fileHash)
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	// Lê a resposta do servidor
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("Server response:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
