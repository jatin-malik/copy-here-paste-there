package main

import (
	"runtime"

	"github.com/jatin-malik/copy-here-paste-there/clipboard"
)

func main() {
	var cboard clipboard.Clipboard

	switch runtime.GOOS {
	case "darwin":
		cboard = clipboard.Clipboard_darwin{}
	default:
		panic("OS not supported")
	}

}
