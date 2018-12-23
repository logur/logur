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
	return &PrintLogger{
		logger: logger,
	}
}

// NewPrintErrorLogger returns a new PrintLogger that logs everything on error level.
func NewPrintErrorLogger(logger Logger) *PrintLogger {
	return NewPrintLogger(LevelFunc(logger, Error))
}

// Print logs a line with fmt.Print semantics.
func (l *PrintLogger) Print(v ...interface{}) {
	l.logger(fmt.Sprint(v...), nil)
}

// Println logs a line with fmt.Println semantics.
func (l *PrintLogger) Println(v ...interface{}) {
	l.logger(strings.TrimSuffix(fmt.Sprintln(v...), "\n"), nil)
}

// Printf logs a line with fmt.Printf semantics.
func (l *PrintLogger) Printf(format string, args ...interface{}) {
	l.logger(fmt.Sprintf(format, args...), nil)
}
