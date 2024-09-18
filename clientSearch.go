package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

type FileSumResult struct {
    Sum int
}

func main() {
    // Endereço do servidor TCP
    address := "150.165.74.30:8000"

    // Conectando ao servidor TCP
    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Println("Erro ao conectar:", err)
        return
    }

	handleConn(conn)
}

func handleConn(conn net.Conn) {
	defer conn.Close()

    if len(os.Args) < 2 {
        fmt.Println("Uso: go run clientSearch.go <numero>")
        return
    }

    // Acessando o argumento
    arg := os.Args[1]

    // Calculando a soma dos arquivos
    scanner := bufio.NewScanner(conn)
    _, err := conn.Write([]byte("search " + arg + "\n"))
	if err != nil {
		fmt.Println("Erro ao enviar mensagem:", err)
		return
	}
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    // Verificando se houve erro na leitura
    if err := scanner.Err(); err != nil {
        fmt.Println("Erro ao ler resposta:", err)
    }
}
