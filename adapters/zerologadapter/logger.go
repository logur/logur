// Package zerologadapter provides a logur adapter for zerolog.
package zerologadapter

import (
	"github.com/rs/zerolog"
)

// Logger is a logur adapter for zerolog.
type Logger struct {
	logger zerolog.Logger
}

// New returns a new logur compatible logger with zerolog as the logging library.
func New(logger zerolog.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Trace implements the logur.Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	// Fall back to Debug
	l.Debug(msg, fields...)
}

// Debug implements the logur.Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	event := l.logger.Debug()

	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}

// Info implements the logur.Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	event := l.logger.Info()

	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}

// Warn implements the logur.Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	event := l.logger.Warn()

	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}

// Error implements the logur.Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	event := l.logger.Error()

	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}
