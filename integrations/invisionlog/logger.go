/*
Package invisionlog provides a github.com/InVisionApp/go-logger logger.

With logur you can easily wire the logging library of your choice into any InVision logger compatible library:

	package main

	import (
		"github.com/goph/logur"
		"github.com/goph/logur/integrations/invisionlog"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		ilogger := invisionlog.New(logger)

		// inject the logger somewhere
	}
*/
package invisionlog

import (
	"fmt"
	"strings"

	"github.com/InVisionApp/go-logger"
	"github.com/goph/logur"
)

type logger struct {
	logger logur.Logger
}

// New returns a new github.com/InVisionApp/go-logger.Logger logger.
func New(l logur.Logger) log.Logger {
	return &logger{l}
}

func (l *logger) Debug(msg ...interface{}) {
	l.logger.Debug(fmt.Sprint(msg...), nil)
}

func (l *logger) Info(msg ...interface{}) {
	l.logger.Info(fmt.Sprint(msg...), nil)
}

func (l *logger) Warn(msg ...interface{}) {
	l.logger.Warn(fmt.Sprint(msg...), nil)
}

func (l *logger) Error(msg ...interface{}) {
	l.logger.Error(fmt.Sprint(msg...), nil)
}

func (l *logger) Debugln(msg ...interface{}) {
	l.Debug(strings.TrimSuffix(fmt.Sprintln(msg...), "\n"))
}

func (l *logger) Infoln(msg ...interface{}) {
	l.Info(strings.TrimSuffix(fmt.Sprintln(msg...), "\n"))
}

func (l *logger) Warnln(msg ...interface{}) {
	l.Warn(strings.TrimSuffix(fmt.Sprintln(msg...), "\n"))
}

func (l *logger) Errorln(msg ...interface{}) {
	l.Error(strings.TrimSuffix(fmt.Sprintln(msg...), "\n"))
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
	return &logger{logur.WithFields(l.logger, fields)}
}
