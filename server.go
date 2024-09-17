package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

// Estrutura para armazenar informações do cliente
type Client struct {
	conn  net.Conn
	canal chan string
}

// Lista de clientes conectados
var clients = make(map[net.Conn]*Client)
var register = make(map[string][]string)
var mu sync.Mutex

func main() {
	listener, err := net.Listen("tcp", "150.165.42.160:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		client := &Client{
			conn:  conn,
			canal: make(chan string),
		}

		// Adiciona o cliente à lista de clientes
		mu.Lock()
		clients[conn] = client
		mu.Unlock()

		go handleConn(client)
	}
}

func handleConn(client *Client) {
	defer func() {
		client.conn.Close()

		// Remove o cliente da lista de clientes
		mu.Lock()
		delete(clients, client.conn)
		mu.Unlock()
	}()

	remoteAddr := client.conn.RemoteAddr().String()
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Println("Erro ao obter o IP:", err)
		return
	}
	fmt.Printf("Nova conexão de %s\n", host)
	appendToRegister(host, "")

	scanner := bufio.NewScanner(client.conn)
	for scanner.Scan() {
		message := scanner.Text()
		appendToRegister(host, message)
		printRegister()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler a mensagem:", err)
	}
}

func appendToRegister(ip string, filename string) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := register[ip]; !exists {
		// Se o IP não existir, cria uma nova entrada com uma lista vazia
		register[ip] = []string{}
	}

	if filename != "" {
		// Adiciona o filename à lista se não for uma string vazia
		register[ip] = append(register[ip], filename)
	}
}

// Obtém a lista de arquivos para um IP específico
func getFiles(ip string) []string {
	mu.Lock()
	defer mu.Unlock()
	return register[ip]
}

// Mostra o conteúdo do registro
func printRegister() {
	mu.Lock()
	defer mu.Unlock()

	for ip, files := range register {
		fmt.Printf("IP: %s\n", ip)
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}
}
