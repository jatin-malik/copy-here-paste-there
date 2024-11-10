package wire

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"

	"github.com/jatin-malik/copy-here-paste-there/config"
)

func encodeToWire(message string) ([]byte, error) {
	// Uses length prefixed messages

	messageBytes := []byte(message)
	messageLength := int32(len(messageBytes))

	buf := new(bytes.Buffer)

	// Write the length of the message as a 4-byte header (big-endian)
	err := binary.Write(buf, binary.BigEndian, messageLength)
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(messageBytes)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decodeFromWire(conn net.Conn) (string, error) {

	var messageLength int32
	err := binary.Read(conn, binary.BigEndian, &messageLength)
	if err != nil {
		return "", err
	}

	messageBytes := make([]byte, messageLength)

	_, err = io.ReadFull(conn, messageBytes)
	if err != nil {
		return "", err
	}

	return string(messageBytes), nil
}

func WriteToConnection(conn net.Conn, to string) {
	var prev_output string
	for {
		output := config.Cboard.Read()
		if output != prev_output {

			// state changed, send it

			message, err := encodeToWire(output)
			if err != nil {
				log.Println("error while encoding message:", err)
				continue
			}
			log.Printf("Sending to %s:%s", to, string(message))
			if _, err := conn.Write(message); err != nil {
				log.Fatal(err)
			}
			prev_output = output
		}

		time.Sleep(2 * time.Second)

	}
}

func ReadFromConnection(conn net.Conn, from string) {
	for {
		message, err := decodeFromWire(conn)
		if err != nil {
			log.Println("Error reading from server:", err)
			return
		}

		log.Printf("Received from %s:%s", from, message)

		// Write to system clipboard
		config.Cboard.Write(message)
	}
}
