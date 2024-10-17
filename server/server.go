package server

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"strings"
)

type Server struct {
	Address string
}

func (s *Server) NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

func (s *Server) StartServer() {
	// Listen on port 6379
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		slog.Error("Could not listen on given port", "err", err)
		return
	}
	defer ln.Close()
	fmt.Printf("Server is listening on port %s...", s.Address)

	// Start accepting connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Error("Could not accept connection", "err", err)
			continue
		}

		// Handle the connection in a goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			slog.Error("Could not read data from connection", "err", err)
			return
		}

		message := strings.TrimSpace(data)
		fmt.Println("Received message: %s", data)

		var response string
		if message == "*1$4PING" {
			response = "+PONG\r\n"
		} else {
			response = "+OK\r\n"
		}

		_, err = conn.Write([]byte(response))
		if err != nil {
			slog.Error("Could not write to connection", "err", err)
			return
		}
	}
}
