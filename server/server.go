package server

import (
	"fmt"
	"net"

	"github.com/jatin-malik/copy-here-paste-there/wire"
)

func Start(port int) {

	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// One connection at a time
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	go wire.ReadFromConnection(conn, "client")
	wire.WriteToConnection(conn, "client")
}
