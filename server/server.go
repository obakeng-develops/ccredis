package server

import (
	"log/slog"
	"net"
)

type Server struct {
	address string
}

func (s *Server) NewServer(address string) *Server {
	return &Server{
		address: address,
	}
}

func (s *Server) StartServer() {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		slog.Error("An error occurred", "err", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Error("An error occurred", "err", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
}
