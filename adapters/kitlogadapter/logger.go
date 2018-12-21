// Package kitlogadapter provides a logur adapter for go-kit logger.
package kitlogadapter

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/logur/internal/keyvals"
)

// Logger is a logur adapter for go-kit logger.
type Logger struct {
	logger log.Logger
}

// New returns a new logur compatible logger with go-kit log as the logging library.
// If nil is passed as logger, a noop logger is used as fallback.
func New(logger log.Logger) *Logger {
	if logger == nil {
		logger = log.NewNopLogger()
	}

	return &Logger{
		logger: logger,
	}
}

// Trace implements the logur.Logger interface.
func (l *Logger) Trace(msg string, fields map[string]interface{}) {
	// Fall back to Debug
	l.Debug(msg, fields)
}

// Debug implements the logur.Logger interface.
func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	_ = level.Debug(l.logger).Log(append(keyvals.FromMap(fields), "msg", msg)...)
}

// Info implements the logur.Logger interface.
func (l *Logger) Info(msg string, fields map[string]interface{}) {
	_ = level.Info(l.logger).Log(append(keyvals.FromMap(fields), "msg", msg)...)
}

// Warn implements the logur.Logger interface.
func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	_ = level.Warn(l.logger).Log(append(keyvals.FromMap(fields), "msg", msg)...)
}

// Error implements the logur.Logger interface.
func (l *Logger) Error(msg string, fields map[string]interface{}) {
	_ = level.Error(l.logger).Log(append(keyvals.FromMap(fields), "msg", msg)...)
}
