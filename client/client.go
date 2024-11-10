package client

import (
	"fmt"
	"log"
	"net"

	"github.com/jatin-malik/copy-here-paste-there/wire"
)

func Start(host string, port int) {
	log.Printf("Connecting to tcp server at (%s:%d)\n", host, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	go wire.ReadFromConnection(conn, "server")
	wire.WriteToConnection(conn, "server")
}
