package logwrapper

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var DefaultLogLevel int
var loglevels map[string]int

// Available log handlers
var (
	TraceLogger   *log.Logger
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {

	// Create a map of the available levels
	loglevels = map[string]int{
		"NOTSET":   0,
		"DEBUG":    10,
		"INFO":     20,
		"WARNING":  30,
		"ERROR":    40,
		"CRITICAL": 50,
	}

	// Set default log level to 0/NOTSET
	DefaultLogLevel = loglevels["NOTSET"]

	// All log handlers print to stdout by default
	TraceLogger = log.New(os.Stdout, "TRACE ", log.Ldate|log.Ltime)
	DebugLogger = log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime)
	InfoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime)

	// Error logger prints to stderr by default
	ErrorLogger = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime)
}

// SetLogLevel sets the default log level.
func SetLogLevel(level string) error {
	value, result := loglevels[level]

	if result == false {
		message := fmt.Sprintf("INVALID_LOG_LEVEL (`%s` does not exist)", level)
		return errors.New(message)
	}

	DefaultLogLevel = value

	return nil
}

// Trace is a wrapper around the `log.Println` method.
func Trace(messages ...interface{}) {
	TraceLogger.Println(messages...)
}

// Debug is a wrapper around the `log.Println` method.
func Debug(messages ...interface{}) {

	if DefaultLogLevel > 10 {
		return
	}

	DebugLogger.Println(messages...)
}

// Info is a wrapper around the `log.Println` method.
func Info(messages ...interface{}) {

	if DefaultLogLevel > 20 {
		return
	}

	InfoLogger.Println(messages...)
}

// Warning is a wrapper around the `log.Println` method.
func Warning(messages ...interface{}) {

	if DefaultLogLevel > 30 {
		return
	}
	WarningLogger.Println(messages...)
}

// Error is a wrapper around the `log.Println` method.
func Error(messages ...interface{}) {
	ErrorLogger.Println(messages...)
}

// Fatal is a wrapper around the `log.Fatal` method which will exit the application.
func Fatal(messages ...interface{}) {
	ErrorLogger.Fatal(messages...)
}

// SetOutput is a wrapper around the `log.Setoutput` method for any io.Writer type
func SetOutput(output io.Writer) {
	TraceLogger.SetOutput(output)
	DebugLogger.SetOutput(output)
	InfoLogger.SetOutput(output)
	WarningLogger.SetOutput(output)
	ErrorLogger.SetOutput(output)
}
