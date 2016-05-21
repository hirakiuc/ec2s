package common

import (
	"github.com/mgutz/ansi"
)

type TextFormatter struct {
	Colored bool
}

func colorForLogLevel(level LogLevel) string {
	switch level {
	case PanicLevel:
		return "red"
	case FatalLevel:
		return "red"
	case ErrorLevel:
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

func (f *TextFormatter) Format(level LogLevel, message string) ([]byte, error) {
	if f.Colored == false {
		return []byte(message), nil
	}

	msg := ansi.Color(message, colorForLogLevel(level))
	return []byte(msg), nil
}
