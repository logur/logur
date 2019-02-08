/*
Package watermilllog provides a Watermill logger.

Watermill accepts a logger that implements the github.com/ThreeDotsLabs/watermill.LoggerAdapter interface.

With logur you can easily wire the logging library of your choice into Watermill:

	package main

	import (
		"github.com/goph/logur"
		"github.com/goph/logur/integrations/watermilllog"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		wlogger := watermilllog.New(logger)

		// inject the logger somewhere
	}
*/
package watermilllog

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/goph/logur"
)

type Logger struct {
	logger logur.Logger
}

// New returns a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
func New(l logur.Logger) *Logger {
	return &Logger{l}
}

func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	fields["err"] = err

	l.logger.Error(msg, fields)
}

func (l *Logger) Info(msg string, fields watermill.LogFields) {
	l.logger.Info(msg, fields)
}

func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	l.logger.Debug(msg, fields)
}

func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	l.logger.Trace(msg, fields)
}

func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	if len(fields) == 0 {
		return l
	}

	return New(logur.WithFields(l.logger, fields))
}
