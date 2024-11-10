package main

import (
	"flag"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/jatin-malik/copy-here-paste-there/client"
	"github.com/jatin-malik/copy-here-paste-there/clipboard"
	"github.com/jatin-malik/copy-here-paste-there/config"
	"github.com/jatin-malik/copy-here-paste-there/server"
)

var modeFlag = flag.String("mode", "server", "{server|client}")
var logLevelFlag = flag.String("log", "info", "log level for the app {debug|info|warn|erorr}")
var hostFlag = flag.String("host", "localhost", "host for remote tcp server")
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

	// Configure logger
	logLevel := parseLogLevel(*logLevelFlag)

	config.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(config.Logger)

	if strings.ToLower(*modeFlag) == "client" {
		// client mode
		client.Start(*hostFlag, *portFlag)
	} else {
		// server mode
		server.Start(*localPortFlag)
	}

}

func parseLogLevel(level string) (logLevel slog.Level) {
	switch strings.ToUpper(level) {
	case "DEBUG":
		logLevel = slog.LevelDebug
	case "INFO":
		logLevel = slog.LevelInfo
	case "WARN":
		logLevel = slog.LevelWarn
	case "ERROR":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	return
}
