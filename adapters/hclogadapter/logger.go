// Package hclogadapter provides a logur compatible adapter for hclog.
package hclogadapter

import (
	"github.com/goph/logur"
	"github.com/goph/logur/internal/keyvals"
	"github.com/hashicorp/go-hclog"
)

// Logger is a logur compatible logger for hclog.
type Logger struct {
	logger hclog.Logger
}

// New returns a new logur compatible logger with hclog as the logging library.
// If nil is passed as logger, the global hclog instance is used as fallback.
func New(logger hclog.Logger) *Logger {
	if logger == nil {
		logger = hclog.Default()
	}

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Trace(msg string, fields map[string]interface{}) {
	if !l.logger.IsTrace() {
		return
	}

	l.logger.Trace(msg, keyvals.FromMap(fields)...)
}

func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	if !l.logger.IsDebug() {
		return
	}

	l.logger.Debug(msg, keyvals.FromMap(fields)...)
}

func (l *Logger) Info(msg string, fields map[string]interface{}) {
	if !l.logger.IsInfo() {
		return
	}

	l.logger.Info(msg, keyvals.FromMap(fields)...)
}

func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	if !l.logger.IsWarn() {
		return
	}

	l.logger.Warn(msg, keyvals.FromMap(fields)...)
}

func (l *Logger) Error(msg string, fields map[string]interface{}) {
	if !l.logger.IsError() {
		return
	}

	l.logger.Error(msg, keyvals.FromMap(fields)...)
}

func (l *Logger) LevelEnabled(level logur.Level) bool {
	switch level {
	case logur.Trace:
		return l.logger.IsTrace()
	case logur.Debug:
		return l.logger.IsDebug()
	case logur.Info:
		return l.logger.IsInfo()
	case logur.Warn:
		return l.logger.IsWarn()
	case logur.Error:
		return l.logger.IsError()
	}

	return true
}
