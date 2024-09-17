package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type FileSumResult struct {
    FileName string
    Sum      int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Err in entry")
		return
	}
	go conection(os.Args[1])
	results := sumFilesInDirectory(os.Args[2])
	fmt.Println(results)
}

func conection(port string) {
	listener, err := net.Listen("tcp", "localhost:"+port)

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
		switch command {
			case "search":
				fmt.Println("A vida é bela") 
		}
	}
}

func sumFilesInDirectory(dir string) ([]FileSumResult) {
	files, _ := ioutil.ReadDir(dir)
	
	if len(files) == 0 {
		fmt.Println("No files found in directory.")
		return nil
	}

	var results []FileSumResult

	for _, file := range files{
		filePath := filepath.Join(dir, file.Name())
		result, _ := sum(filePath)
		results = append(results, FileSumResult{
			FileName: file.Name(),
			Sum:      result,
		})
	}

	return results
}

func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
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