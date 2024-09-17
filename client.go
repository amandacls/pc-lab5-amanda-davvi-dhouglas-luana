package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
    "path/filepath"
)

type FileSumResult struct {
    Sum int `json:"sum"`
}

func main() {
    // Endereço do servidor TCP
    address := "150.165.42.160:8000"

    // Conectando ao servidor TCP
    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Println("Erro ao conectar:", err)
        return
    }
    defer conn.Close()

    // Calculando a soma dos arquivos
    results := sumFilesInDirectory("/tmp/dataset")

    // Serializando os resultados para JSON
    resultsJSON, err := json.Marshal(results)
    if err != nil {
        fmt.Println("Erro ao serializar resultados:", err)
        return
    }

    // Enviando uma mensagem
    _, err = conn.Write([]byte("Oi servidor\n"))
    if err != nil {
        fmt.Println("Erro ao enviar mensagem:", err)
        return
    }

    // Enviando a lista de resultados como JSON
    _, err = conn.Write(resultsJSON)
    if err != nil {
        fmt.Println("Erro ao enviar dados:", err)
        return
    }
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