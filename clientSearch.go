package main

import (
    "fmt"
    "net"
    "bufio"
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

    // Calculando a soma dos arquivos
    scanner := bufio.NewScanner(conn)
    _, err := conn.Write([]byte("search 1336472306\n"))
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
