package clipboard

import (
	"os/exec"
	"strings"
)

type Clipboard_darwin struct {
}

func (cd Clipboard_darwin) Read() string {
	// reads from the macOS system clipboard
	cmd := exec.Command("pbpaste")
	output, err := cmd.Output()
	if err != nil {
		// ignoring error handling intentionally
		return ""
	}
	return string(output)
}

func (cd Clipboard_darwin) Write(text string) {
	// writes to the macOS system clipboard
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	cmd.Run() // ignoring error handling intentionally
}
