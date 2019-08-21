// Package logrusadapter provides a logur adapter for logrus.
package logrusadapter

import (
	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
)

// Logger is a logur adapter for logrus.
// Deprecated: use logur.dev/adapter/logrus.Logger instead.
type Logger = logrusadapter.Logger

// New returns a new logur compatible logger with logrus as the logging library.
// If nil is passed as logger, the global logrus instance is used as fallback.
// Deprecated: use logur.dev/adapter/logrus.New instead.
func New(logger *logrus.Logger) *Logger {
	return logrusadapter.New(logger)
}

// NewFromEntry returns a new logur compatible logger with logrus as the
// logging library while preserving pre-set fields.
// If nil is passed as entry, the global logrus instance is used as fallback.
// Deprecated: use logur.dev/adapter/logrus.NewFromEntry instead.
func NewFromEntry(entry *logrus.Entry) *Logger {
	return logrusadapter.NewFromEntry(entry)
}
