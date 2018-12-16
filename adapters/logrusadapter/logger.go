// Package logrusadapter provides a logur compatible adapter for logrus.
package logrusadapter

import (
	"github.com/goph/logur"
	"github.com/sirupsen/logrus"
)

type adapter struct {
	*logrus.Entry
}

// New returns a new logur compatible logger with Logrus as the logging library.
// If nil is passed as logger, the global Logrus instance is used as fallback.
func New(logger *logrus.Logger) logur.Logger {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	return &adapter{logrus.NewEntry(logger)}
}

// WithFields returns a new logger based on the original logger with
// the additional supplied fields.
func (a *adapter) WithFields(fields map[string]interface{}) logur.Logger {
	return &adapter{
		a.Entry.WithFields(logrus.Fields(fields)),
	}
}
