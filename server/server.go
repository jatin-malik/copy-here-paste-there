package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/jatin-malik/copy-here-paste-there/config"
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

	go readFromClient(conn)
	writeToClient(conn)
}

func readFromClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", err)
			return
		}
		message = strings.TrimSpace(message)
		fmt.Println("Received from client:", message)

		// Write to system clipboard
		config.Cboard.Write(message)
	}

}

func writeToClient(conn net.Conn) {
	var prev_output string
	for {
		output := config.Cboard.Read()
		if output != prev_output {
			// state changed, send it
			fmt.Println("Sending to client:", output)
			if _, err := conn.Write([]byte(output + "\n")); err != nil {
				log.Fatal(err)
			}
			prev_output = output
		}

		time.Sleep(2 * time.Second)

	}
}
