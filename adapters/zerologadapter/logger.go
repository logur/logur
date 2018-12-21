// Package zerologadapter provides a logur compatible adapter for zerolog.
package zerologadapter

import (
	"github.com/rs/zerolog"
)

// Logger is a logur compatible logger for zerolog.
type Logger struct {
	logger zerolog.Logger
}

// New returns a new logur compatible logger with zerolog as the logging library.
func New(logger zerolog.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Trace(msg string, fields map[string]interface{}) {
	// Fall back to Debug
	l.Debug(msg, fields)
}

func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	l.logger.Debug().Fields(fields).Msg(msg)
}

func (l *Logger) Info(msg string, fields map[string]interface{}) {
	l.logger.Info().Fields(fields).Msg(msg)
}

func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	l.logger.Warn().Fields(fields).Msg(msg)
}

func (l *Logger) Error(msg string, fields map[string]interface{}) {
	l.logger.Error().Fields(fields).Msg(msg)
}
