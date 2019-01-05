// Package zapadapter provides a logur adapter for Uber's Zap.
package zapadapter

import (
	"github.com/goph/logur"
	"github.com/goph/logur/internal/keyvals"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a logur adapter for Uber's zap.
type Logger struct {
	logger *zap.SugaredLogger
	core   zapcore.Core
}

// New returns a new logur compatible logger with Uber's zap as the logging library.
// If nil is passed as logger, the global sugared logger instance is used as fallback.
func New(logger *zap.Logger) *Logger {
	if logger == nil {
		logger = zap.L()
	}

	return &Logger{
		logger: logger.Sugar(),
		core:   logger.Core(),
	}
}

// Trace implements the logur.Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	// Fall back to Debug
	l.Debug(msg, fields...)
}

// Debug implements the logur.Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	if !l.core.Enabled(zap.DebugLevel) {
		return
	}

	l.logger.Debugw(msg, l.keyvals(fields)...)
}

// Info implements the logur.Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	if !l.core.Enabled(zap.InfoLevel) {
		return
	}

	l.logger.Infow(msg, l.keyvals(fields)...)
}

// Warn implements the logur.Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	if !l.core.Enabled(zap.WarnLevel) {
		return
	}

	l.logger.Warnw(msg, l.keyvals(fields)...)
}

// Error implements the logur.Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	if !l.core.Enabled(zap.ErrorLevel) {
		return
	}

	l.logger.Errorw(msg, l.keyvals(fields)...)
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
		return l.core.Enabled(zap.DebugLevel)
	case logur.Debug:
		return l.core.Enabled(zap.DebugLevel)
	case logur.Info:
		return l.core.Enabled(zap.InfoLevel)
	case logur.Warn:
		return l.core.Enabled(zap.WarnLevel)
	case logur.Error:
		return l.core.Enabled(zap.ErrorLevel)
	}

	return true
}
