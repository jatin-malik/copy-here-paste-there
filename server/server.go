package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/jatin-malik/copy-here-paste-there/wire"
)

func Start(port int) {

	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		slog.Error(fmt.Sprintf("error while creating a tcp listener: %s", err))
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Error(fmt.Sprintf("error while accepting connection: %s", err))
			continue
		}
		slog.Debug(fmt.Sprintf("Accepted a connection %v", conn))

		// One connection at a time
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		conn.Close()
		cancel()
	}()

	go wire.WriteToConnection(ctx, conn, "client")
	wire.ReadFromConnection(conn, "client")
}
