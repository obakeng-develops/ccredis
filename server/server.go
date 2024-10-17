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
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		slog.Error("An error occurred", "err", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server is listening on port 6379...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Error("An error occurred", "err", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
}
