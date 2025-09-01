package logger

import (
    "fmt"
    "os"
)

// Logger struct to hold the log file
type Logger struct {
    file *os.File
}

// NewLogger initializes a new logger
func NewLogger(filePath string) (*Logger, error) {
    // Open the log file for appending
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        return nil, err // Return error if file can't be opened
    }

    return &Logger{file: file}, nil // Return the logger instance
}

// Log writes a log message to the log file
func (l *Logger) Log(message string) {
    // Format the log message
    logMessage := fmt.Sprintf("%s: %s\n", getCurrentTime(), message)
    // Write the log message to the file
    l.file.WriteString(logMessage)
}

// Close closes the log file
func (l *Logger) Close() {
    l.file.Close() // Close the file when done
}

// getCurrentTime returns the current time as a string
func getCurrentTime() string {
    return fmt.Sprintf("%v", os.Args) // Placeholder for current time, can be improved
}