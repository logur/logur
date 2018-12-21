// Package logrusadapter provides a logur compatible adapter for logrus.
package logrusadapter

import (
	"github.com/goph/logur"
	"github.com/sirupsen/logrus"
)

// Logger is a logur compatible logger for logrus.
type Logger struct {
	entry *logrus.Entry
}

// New returns a new logur compatible logger with logrus as the logging library.
// If nil is passed as logger, the global logrus instance is used as fallback.
func New(logger *logrus.Logger) *Logger {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	return &Logger{
		entry: logrus.NewEntry(logger),
	}
}

func (l *Logger) Trace(msg string, fields map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.TraceLevel) {
		return
	}

	l.entry.WithFields(logrus.Fields(fields)).Trace(msg)
}

func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.DebugLevel) {
		return
	}

	l.entry.WithFields(logrus.Fields(fields)).Debug(msg)
}

func (l *Logger) Info(msg string, fields map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.InfoLevel) {
		return
	}

	l.entry.WithFields(logrus.Fields(fields)).Info(msg)
}

func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.WarnLevel) {
		return
	}

	l.entry.WithFields(logrus.Fields(fields)).Warn(msg)
}

func (l *Logger) Error(msg string, fields map[string]interface{}) {
	if !l.entry.Logger.IsLevelEnabled(logrus.ErrorLevel) {
		return
	}

	l.entry.WithFields(logrus.Fields(fields)).Error(msg)
}

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
