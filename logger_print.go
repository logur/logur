package logur

import (
	"fmt"
	"strings"
)

// PrintLogger logs messages with fmt.Print* function semantics.
type PrintLogger struct {
	logger LogFunc
}

// NewPrintLogger returns a new PrintLogger.
func NewPrintLogger(logger LogFunc) *PrintLogger {
	return &PrintLogger{logger}
}

// NewErrorPrintLogger returns a new PrintLogger that logs everything on error level.
func NewErrorPrintLogger(logger Logger) *PrintLogger {
	return NewPrintLogger(LevelFunc(logger, Error))
}

// Print logs a line with fmt.Print semantics.
func (l *PrintLogger) Print(v ...interface{}) {
	l.logger(fmt.Sprint(v...))
}

// Println logs a line with fmt.Println semantics.
func (l *PrintLogger) Println(v ...interface{}) {
	l.logger(strings.TrimSuffix(fmt.Sprintln(v...), "\n"))
}

// Printf logs a line with fmt.Printf semantics.
func (l *PrintLogger) Printf(format string, args ...interface{}) {
	l.logger(fmt.Sprintf(format, args...))
}

// MessageLogger simplifies the Logger interface by removing the second context parameter.
// Useful when there is no need for contextual logging.
type MessageLogger struct {
	logger Logger
}

// NewMessageLogger returns a new MessageLogger instance.
func NewMessageLogger(logger Logger) *MessageLogger {
	return &MessageLogger{logger}
}

// Trace logs a Trace level event.
func (l *MessageLogger) Trace(msg string) { l.logger.Trace(msg) }

// Debug logs a Debug level event.
func (l *MessageLogger) Debug(msg string) { l.logger.Debug(msg) }

// Info logs a Info level event.
func (l *MessageLogger) Info(msg string) { l.logger.Info(msg) }

// Warn logs a Warn level event.
func (l *MessageLogger) Warn(msg string) { l.logger.Warn(msg) }

// Error logs a Error level event.
func (l *MessageLogger) Error(msg string) { l.logger.Error(msg) }
