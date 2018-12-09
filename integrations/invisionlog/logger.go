// Package invisionlog provides logur integration for github.com/InVisionApp/go-logger.
package invisionlog

import (
	"github.com/InVisionApp/go-logger"
	"github.com/goph/logur"
)

type logger struct {
	logur.Logger
}

// New returns a new github.com/InVisionApp/go-logger.Logger compatible logger.
func New(l logur.Logger) log.Logger {
	return &logger{l}
}

// WithFields returns a new logger with the additional supplied fields.
func (l *logger) WithFields(fields log.Fields) log.Logger {
	return &logger{l.Logger.WithFields(logur.Fields(fields))}
}
