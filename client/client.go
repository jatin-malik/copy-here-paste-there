package client

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/jatin-malik/copy-here-paste-there/wire"
)

func Start(host string, port int) {
	slog.Info(fmt.Sprintf("Connecting to tcp server at (%s:%d)\n", host, port))
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to dial: %v", err))
		os.Exit(1)
	}
	defer conn.Close()

	go wire.WriteToConnection(context.Background(), conn, "server")
	wire.ReadFromConnection(conn, "server")

}
