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
	"github.com/pkg/errors"
)

// ErrorHandler handles an error passed to the logger.
type ErrorHandler interface {
	// Handle handles an error passed to the logger.
	Handle(err error, fields map[string]interface{})
}

// Logger is a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
type Logger struct {
	logger       logur.Logger
	errorHandler ErrorHandler
}

// New returns a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
func New(logger logur.Logger) *Logger {
	return &Logger{logger: logger}
}

// NewWithErrorHandler returns a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
// Compared to the logger returned by New, this logger sends errors to the error handler.
func NewWithErrorHandler(logger logur.Logger, errorHandler ErrorHandler) *Logger {
	return &Logger{
		logger:       logger,
		errorHandler: errorHandler,
	}
}

func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	if l.errorHandler == nil {
		fields["err"] = err

		l.logger.Error(msg, fields)

		return
	}

	err = errors.WithMessage(err, msg)

	l.errorHandler.Handle(err, fields)
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

	return NewWithErrorHandler(logur.WithFields(l.logger, fields), l.errorHandler)
}
