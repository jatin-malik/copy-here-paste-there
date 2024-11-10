package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/jatin-malik/copy-here-paste-there/config"
)

func Start(host string, port int) {
	log.Printf("Connecting to tcp server at (%s:%d)\n", host, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	go readFromServer(conn)
	writeToServer(conn)
}

func writeToServer(conn net.Conn) {
	var prev_output string
	for {
		output := config.Cboard.Read()
		if output != prev_output {
			// state changed, send it
			fmt.Println("Sending:", output)
			if _, err := conn.Write([]byte(output + "\n")); err != nil {
				log.Fatal(err)
			}
			prev_output = output
		}

		time.Sleep(2 * time.Second)

	}
}

func readFromServer(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Server disconnected:", err)
			return
		}
		message = strings.TrimSpace(message)
		fmt.Println("Received from server:", message)

		// Write to system clipboard
		config.Cboard.Write(message)
	}
}
