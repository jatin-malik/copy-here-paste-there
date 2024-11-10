package config

import (
	"log/slog"

	"github.com/jatin-malik/copy-here-paste-there/clipboard"
)

// Global variables used across the application
var Cboard clipboard.Clipboard
var Logger *slog.Logger
