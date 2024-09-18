package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "path/filepath"
    "encoding/json"
	// "strconv"
    "bufio"
	// "time"
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
	// results := sumFilesInDirectory("/tmp/dataset")
    scanner := bufio.NewScanner(conn)
    _, err := conn.Write([]byte("search 1336472306\n"))
	if err != nil {
		fmt.Println("Erro ao enviar mensagem:", err)
		return
	}
	// encoder := json.NewEncoder(conn)
    // if err := encoder.Encode(results); err != nil {
    //     fmt.Println("Error sending file info:", err)
    // }

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    // Verificando se houve erro na leitura
    if err := scanner.Err(); err != nil {
        fmt.Println("Erro ao ler resposta:", err)
    }

	// for _, result := range(results) {
	// 	r := strconv.Itoa(result.Sum) + "\n"
	// 	// fmt.Print(r)
	// 	_, err := conn.Write([]byte(r))
	// 	if err != nil {
	// 		fmt.Println("Erro ao enviar mensagem:", err)
	// 		return
	// 	}
	// }
	// time.Sleep(1 * time.Minute)
	
}

func sumFilesInDirectory(dir string) []FileSumResult {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Println("Erro ao ler diretório:", err)
        return nil
    }

    if len(files) == 0 {
        fmt.Println("Nenhum arquivo encontrado no diretório.")
        return nil
    }

    var results []FileSumResult

    for _, file := range files {
        filePath := filepath.Join(dir, file.Name())
        result, err := sum(filePath)
        if err != nil {
            fmt.Printf("Erro ao somar arquivo %s: %v\n", filePath, err)
            continue
        }
        results = append(results, FileSumResult{
            Sum: result,
        })
    }

    return results
}

func readFile(filePath string) ([]byte, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Erro ao ler arquivo %s: %v\n", filePath, err)
        return nil, err
    }
    return data, nil
}

func sum(filePath string) (int, error) {
    data, err := readFile(filePath)
    if err != nil {
        return 0, err
    }

    _sum := 0
    for _, b := range data {
        _sum += int(b)
    }

    return _sum, nil
}

func sendFileInfo(conn net.Conn, size FileSumResult) {
    encoder := json.NewEncoder(conn)
    if err := encoder.Encode(size); err != nil {
        fmt.Println("Error sending file info:", err)
    }
}