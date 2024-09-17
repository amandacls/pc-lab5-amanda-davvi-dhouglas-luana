package main


import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Err in entry")
		return
	}
	go conection(os.Args[1])
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