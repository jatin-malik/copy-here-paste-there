package main

import (
	"flag"
	"runtime"
	"strings"

	"github.com/jatin-malik/copy-here-paste-there/client"
	"github.com/jatin-malik/copy-here-paste-there/clipboard"
	"github.com/jatin-malik/copy-here-paste-there/config"
	"github.com/jatin-malik/copy-here-paste-there/server"
)

var hostFlag = flag.String("host", "localhost", "host for remote tcp server")
var modeFlag = flag.String("mode", "client", "{server|client}")
var portFlag = flag.Int("port", 80, "port for remote tcp server")
var localPortFlag = flag.Int("localport", 80, "port for local tcp server")

func init() {

	switch runtime.GOOS {
	case "darwin":
		config.Cboard = clipboard.Clipboard_darwin{}
	default:
		panic("OS not supported")
	}
}

func main() {

	flag.Parse()

	if strings.ToLower(*modeFlag) == "client" {
		// client mode
		client.Start(*hostFlag, *portFlag)
	} else {
		// server mode
		server.Start(*localPortFlag)
	}

}
