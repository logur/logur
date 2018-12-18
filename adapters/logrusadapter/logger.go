// Package logrusadapter provides a logur compatible adapter for logrus.
package logrusadapter

import (
	"github.com/goph/logur"
	"github.com/sirupsen/logrus"
)

type adapter struct {
	logger *logrus.Entry
}

// New returns a new logur compatible logger with Logrus as the logging library.
// If nil is passed as logger, the global Logrus instance is used as fallback.
func New(logger *logrus.Logger) logur.Logger {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	return &adapter{logrus.NewEntry(logger)}
}

func (a *adapter) Trace(msg string, fields map[string]interface{}) {
	a.logger.WithFields(logrus.Fields(fields)).Trace(msg)
}

func (a *adapter) Debug(msg string, fields map[string]interface{}) {
	a.logger.WithFields(logrus.Fields(fields)).Debug(msg)
}

func (a *adapter) Info(msg string, fields map[string]interface{}) {
	a.logger.WithFields(logrus.Fields(fields)).Info(msg)
}

func (a *adapter) Warn(msg string, fields map[string]interface{}) {
	a.logger.WithFields(logrus.Fields(fields)).Warn(msg)
}

func (a *adapter) Error(msg string, fields map[string]interface{}) {
	a.logger.WithFields(logrus.Fields(fields)).Error(msg)
}

// WithFields returns a new logger based on the original logger with
// the additional supplied fields.
func (a *adapter) WithFields(fields map[string]interface{}) logur.Logger {
	return &adapter{
		a.logger.WithFields(logrus.Fields(fields)),
	}
}
