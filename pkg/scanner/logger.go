package scanner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Logger struct {
	mu       sync.Mutex
	logFile  *os.File
	logLevel string
}

// Log levels
const (
	InfoLevel  = "INFO"
	WarnLevel  = "WARN"
	ErrorLevel = "ERROR"
)

// New creates a new Logger instance with optional log file
func New(logFilePath string, logLevel string) (*Logger, error) {
	logger := &Logger{
		logLevel: logLevel,
	}

	// Create the log file if it does not exist
	if logFilePath != "" {
		if err := os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm); err != nil {
			return nil, err
		}
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		logger.logFile = file
	}

	return logger, nil
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	l.log(InfoLevel, msg)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string) {
	l.log(WarnLevel, msg)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.log(ErrorLevel, msg)
}

// log writes the log message to console and file
func (l *Logger) log(level string, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	logMsg := fmt.Sprintf("%s: %s", level, msg)
	fmt.Println(logMsg) // Print to console

	// Log to file if it exists
	if l.logFile != nil {
		log.SetOutput(l.logFile)
		log.Println(logMsg)
	}
}

// Close closes the log file if it was opened
func (l *Logger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}
