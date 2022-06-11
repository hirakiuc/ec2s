package common

import (
	"github.com/mgutz/ansi"
)

// TextFormatter define log formatter.
type TextFormatter struct {
	Colored bool
}

func colorForLogLevel(level LogLevel) string {
	switch level {
	case PanicLevel, FatalLevel, ErrorLevel:
		return "red"
	case WarnLevel:
		return "yellow"
	case InfoLevel:
		return "green"
	case DebugLevel:
		return ansi.DefaultFG
	}

	return ansi.DefaultFG
}

// Format return formatted string.
func (f *TextFormatter) Format(level LogLevel, message string) ([]byte, error) {
	if !f.Colored {
		return []byte(message), nil
	}

	msg := ansi.Color(message, colorForLogLevel(level))

	return []byte(msg), nil
}
