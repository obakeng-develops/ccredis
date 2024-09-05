package server

import (
	"log/slog"
	"net"
)

func StartServer() {
	ln, err := net.Listen("tcp", ":8080")
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
