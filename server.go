package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"strings"
	"encoding/json"
	"strconv"
)

// Estrutura para armazenar informações do cliente
type Client struct {
	conn   net.Conn
	canal  chan string
}

type Result struct {
    Sum int `json:"Sum"`
}

// Lista de clientes conectados
var clients = make(map[net.Conn]*Client)
var register = make(map[string][]string)
var mu sync.Mutex

func main() {
	listener, err := net.Listen("tcp", "150.165.74.30:8000")
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
	canal1 := make(chan []string)


	remoteAddr := client.conn.RemoteAddr().String()
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Println("Erro ao obter o IP:", err)
		return
	}
	fmt.Printf("Nova conexão de %s\n", host)
	appendToRegister(host, "")

	go func() {
		scanner := bufio.NewScanner(client.conn)
			for scanner.Scan() {
				message := scanner.Text()
				parts := strings.Split(message, " ")
		
				switch parts[0] {
					case "search":
						result := findIPsWithValue(parts[1])
						canal1 <- result
					default:
						var results []Result

						// Deserializando o JSON
						err := json.Unmarshal([]byte(message), &results)
						if err != nil {
							fmt.Println("Erro ao deserializar JSON:", err)
							return
						}

						// Iterando sobre os resultados
						for _, result := range results {
							sum := strconv.Itoa(result.Sum)
							appendToRegister(host, sum)
						}
						response := []string{"Hash cadastrados:\n", message}
						canal1 <- response
				}
		
			}
		
			if err := scanner.Err(); err != nil {
				fmt.Println("Erro ao ler a mensagem:", err)
			}
	}()
	
	select {
		case msg1 := <- canal1:
			_, err := client.conn.Write([]byte(strings.Join(msg1, " ") + "\n"))
			if err != nil {
				fmt.Println("Erro ao enviar mensagem:", err)
				return
			}
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

func findIPsWithValue(value string) []string {
	var result []string
	result = append(result, "Ips encontrados: ")

	// Itera sobre o mapa
	for ip, values := range register {
		// Itera sobre a lista de valores para cada IP
		for _, v := range values {
			if v == value {
				result = append(result, ip)
				break
			}
		}
	}

	return result
} 