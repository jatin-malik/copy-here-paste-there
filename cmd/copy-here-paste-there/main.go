package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"time"

	"github.com/jatin-malik/copy-here-paste-there/clipboard"
)

var cboard clipboard.Clipboard

var hostFlag = flag.String("host", "localhost", "host for remote tcp server")
var modeFlag = flag.String("mode", "client", "{server|client}")
var portFlag = flag.Int("port", 80, "port for remote tcp server")
var localPortFlag = flag.Int("localport", 80, "port for local tcp server")

func init() {

	switch runtime.GOOS {
	case "darwin":
		cboard = clipboard.Clipboard_darwin{}
	default:
		panic("OS not supported")
	}
}

func clipboard_client(conn net.Conn) {
	var prev_output string
	for {
		output := cboard.Read()
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

func clipboard_server(port int) {

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
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		message = strings.TrimSpace(message)

		fmt.Println("Writing to clipboard:", message)
		cboard.Write(message)
	}
}

func main() {

	flag.Parse()

	if strings.ToLower(*modeFlag) == "client" {
		log.Printf("Connecting to tcp server at (%s:%d)\n", *hostFlag, *portFlag)
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *hostFlag, *portFlag))
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}
		defer conn.Close()

		clipboard_client(conn)
	} else {
		// server mode
		clipboard_server(*localPortFlag)
	}

}
