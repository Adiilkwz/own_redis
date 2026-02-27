package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 8080, "Port to connect to")
	flag.Parse()

	address := fmt.Sprintf("127.0.0.1:%d", *port)

	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to Own Redis on %s! Type 'exit' to quit.\n", address)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}
		if input == "" {
			continue
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Failed to send:", err)
			continue
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Failed to read response:", err)
			continue
		}

		fmt.Print(string(buffer[:n]))
	}
}
