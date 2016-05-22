package common

// This code created with reference sirupsen/logrus
// https://github.com/Sirupsen/logrus/blob/master/entry.go

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

// LogLevel define log level.
type LogLevel uint8

const (
	// PanicLevel is a log level on critical case.
	PanicLevel LogLevel = iota

	// FatalLevel is a log level on fatal case.
	FatalLevel

	// ErrorLevel is a log level on error case.
	ErrorLevel

	// WarnLevel is a log level on warn case.
	WarnLevel

	// InfoLevel is a log level on info logging.
	InfoLevel

	// DebugLevel is a log level on debug logging.
	DebugLevel
)

// Logger define a logger struct.
type Logger struct {
	LogLevel  LogLevel
	Output    io.Writer
	Formatter *TextFormatter
	mu        sync.Mutex
}

var logger = newLogger()

func newLogger() *Logger {
	return &Logger{
		LogLevel: DebugLevel,
		Output:   os.Stdout,
		Formatter: &TextFormatter{
			Colored: true,
		},
	}
}

// GetLogger return Logger instance.
func GetLogger() *Logger {
	return logger
}

// SetColored configure colorize log message or not.
func (log *Logger) SetColored(flag bool) {
	log.Formatter.Colored = flag
}

func (log *Logger) reader(level LogLevel, message string) *bytes.Buffer {
	msg, _ := log.Formatter.Format(level, message)
	return bytes.NewBuffer([]byte(msg))
}

func (log *Logger) log(logLevel LogLevel, msg string) {
	log.mu.Lock()
	defer log.mu.Unlock()

	_, err := io.Copy(log.Output, log.reader(logLevel, msg))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
	}

	if logLevel <= PanicLevel {
		panic(msg)
	}
}

// Debug puts debug level message to log.
func (log *Logger) Debug(format string, args ...interface{}) {
	if log.LogLevel >= DebugLevel {
		log.log(DebugLevel, fmt.Sprintf(format, args...))
	}
}

// Print is a alias to Info.
func (log *Logger) Print(format string, args ...interface{}) {
	log.Info(format, args...)
}

// Info puts info level message to log.
func (log *Logger) Info(format string, args ...interface{}) {
	if log.LogLevel >= InfoLevel {
		log.log(InfoLevel, fmt.Sprintf(format, args...))
	}
}

// Warn puts warn level message to log.
func (log *Logger) Warn(format string, args ...interface{}) {
	if log.LogLevel >= WarnLevel {
		log.log(WarnLevel, fmt.Sprintf(format, args...))
	}
}

// Warning is a alias to Warn.
func (log *Logger) Warning(format string, args ...interface{}) {
	log.Warn(format, args...)
}

// Error puts error message to log.
func (log *Logger) Error(format string, args ...interface{}) {
	if log.LogLevel >= ErrorLevel {
		log.log(ErrorLevel, fmt.Sprintf(format, args...))
	}
}

// Fatal puts fatal message to log.
func (log *Logger) Fatal(format string, args ...interface{}) {
	if log.LogLevel >= FatalLevel {
		log.log(FatalLevel, fmt.Sprintf(format, args...))
	}

	os.Exit(1)
}

// Panic puts panic message to log.
func (log *Logger) Panic(format string, args ...interface{}) {
	if log.LogLevel >= PanicLevel {
		log.log(PanicLevel, fmt.Sprintf(format, args...))
	}
	panic(fmt.Sprint(args...))
}
