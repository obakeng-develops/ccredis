package server

import (
	"fmt"
	"log/slog"
	"net"
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
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		slog.Error("Could not listen on given port", "err", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server is listening on port 6379...")

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

	fmt.Println("Hello")
}
