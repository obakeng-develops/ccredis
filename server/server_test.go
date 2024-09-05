package server_test

import (
	"net"
	"testing"
	"time"

	"github.com/obakeng-develops/redis-server/server"
)

func TestStartServer(t *testing.T) {
	go server.StartServer()

	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	t.Log("Successfully connected to the server")
}
