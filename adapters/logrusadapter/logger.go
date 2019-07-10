// Package logrusadapter provides a logur adapter for logrus.
package logrusadapter

import (
	"github.com/sirupsen/logrus"

	"github.com/goph/logur"
)

// Logger is a logur adapter for logrus.
type Logger struct {
	entry *logrus.Entry
}

// New returns a new logur compatible logger with logrus as the logging library.
// If nil is passed as logger, the global logrus instance is used as fallback.
func New(logger *logrus.Logger) *Logger {
	if logger == nil {
		return NewFromEntry(nil)
	}

	return NewFromEntry(logrus.NewEntry(logger))
}

// NewFromEntry returns a new logur compatible logger with logrus as the
// logging library while preserving pre-set fields.
// If nil is passed as entry, the global logrus instance is used as fallback.
func NewFromEntry(entry *logrus.Entry) *Logger {
	if entry == nil {
		entry = logrus.NewEntry(logrus.StandardLogger())
	}

	return &Logger{
		entry: entry,
	}
}

// Trace implements the logur.Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.TraceLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}

	entry.Trace(msg)
}

// Debug implements the logur.Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.DebugLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}

	entry.Debug(msg)
}

// Info implements the logur.Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.InfoLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}

	entry.Info(msg)
}

// Warn implements the logur.Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.WarnLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}

	entry.Warn(msg)
}

// Error implements the logur.Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.ErrorLevel) {
		return
	}

	var entry = l.entry
	if len(fields) > 0 {
		entry = entry.WithFields(logrus.Fields(fields[0]))
	}

	entry.Error(msg)
}

// LevelEnabled implements logur.LevelEnabler interface.
func (l *Logger) LevelEnabled(level logur.Level) bool {
	switch level {
	case logur.Trace:
		return l.entry.Logger.IsLevelEnabled(logrus.TraceLevel)
	case logur.Debug:
		return l.entry.Logger.IsLevelEnabled(logrus.DebugLevel)
	case logur.Info:
		return l.entry.Logger.IsLevelEnabled(logrus.InfoLevel)
	case logur.Warn:
		return l.entry.Logger.IsLevelEnabled(logrus.WarnLevel)
	case logur.Error:
		return l.entry.Logger.IsLevelEnabled(logrus.ErrorLevel)
	}

	return true
}
