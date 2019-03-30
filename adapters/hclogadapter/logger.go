// Package hclogadapter provides a logur adapter for hclog.
package hclogadapter

import (
	"github.com/hashicorp/go-hclog"

	"github.com/goph/logur"
	"github.com/goph/logur/internal/keyvals"
)

// Logger is a logur adapter for hclog.
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

// Trace implements the logur.Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsTrace() {
		return
	}

	l.logger.Trace(msg, l.keyvals(fields)...)
}

// Debug implements the logur.Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsDebug() {
		return
	}

	l.logger.Debug(msg, l.keyvals(fields)...)
}

// Info implements the logur.Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsInfo() {
		return
	}

	l.logger.Info(msg, l.keyvals(fields)...)
}

// Warn implements the logur.Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsWarn() {
		return
	}

	l.logger.Warn(msg, l.keyvals(fields)...)
}

// Error implements the logur.Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	if !l.logger.IsError() {
		return
	}

	l.logger.Error(msg, l.keyvals(fields)...)
}

func (l *Logger) keyvals(fields []map[string]interface{}) []interface{} {
	var kvs []interface{}
	if len(fields) > 0 {
		kvs = keyvals.FromMap(fields[0])
	}

	return kvs
}

// LevelEnabled implements logur.LevelEnabler interface.
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
