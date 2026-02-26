package server

import (
	"fmt"
	"net"

	"own-redis/internal/protocol"
)

type UDPServer struct {
	Port    int
	Handler *protocol.Handler
}

func NewUDPServer(port int, handler *protocol.Handler) *UDPServer {
	return &UDPServer{
		Port:    port,
		Handler: handler,
	}
}

func (s *UDPServer) Start() error {
	address := fmt.Sprintf("0.0.0.0:%d", s.Port)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return fmt.Errorf("error resolving UDP address: %w", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return fmt.Errorf("error starting UDP server: %w", err)
	}
	defer conn.Close()

	fmt.Printf("Own Redis server started on %s\n", address)

	buffer := make([]byte, 2048)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading from UDP: %s\n", err)
			continue
		}

		request := string(buffer[:n])

		go func(req string, addr *net.UDPAddr) {
			response := s.Handler.HandleProcess(req)
			if response != "" {
				_, err := conn.WriteToUDP([]byte(response), addr)
				if err != nil {
					fmt.Printf("Error writing to UDP: %v\n", err)
				}
			}
		}(request, clientAddr)
	}
}
