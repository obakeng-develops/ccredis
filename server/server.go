package server

import (
	"log/slog"
	"net"
)

type Server struct {
	network string
	address string
}

func (s *Server) NewServer(network, address string) *Server {
	return &Server{
		network: network,
		address: address,
	}
}

func (s *Server) StartServer() {
	ln, err := net.Listen(s.network, s.address)
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
