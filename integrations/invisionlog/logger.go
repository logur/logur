// Package invisionlog provides logur integration for github.com/InVisionApp/go-logger.
package invisionlog

import (
	"fmt"

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

func (l *logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

// WithFields returns a new logger with the additional supplied fields.
func (l *logger) WithFields(fields log.Fields) log.Logger {
	return &logger{l.Logger.WithFields(logur.Fields(fields))}
}
