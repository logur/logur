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
	watermillintegration "logur.dev/integration/watermill"

	"github.com/goph/logur"
)

// ErrorHandler handles an error passed to the logger.
// Deprecated: use logur.dev/integration/watermill.ErrorHandler instead.
type ErrorHandler = watermillintegration.ErrorHandler

// Logger is a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
// Deprecated: use logur.dev/integration/watermill.Logger instead.
type Logger = watermillintegration.Logger

// New returns a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
// Deprecated: use logur.dev/integration/watermill.New instead.
func New(logger logur.Logger) *Logger {
	return watermillintegration.New(logger)
}

// NewWithErrorHandler returns a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
// Compared to the logger returned by New, this logger sends errors to the error handler.
// Deprecated: use logur.dev/integration/watermill.NewWithErrorHandler instead.
func NewWithErrorHandler(logger logur.Logger, errorHandler ErrorHandler) *Logger {
	return watermillintegration.NewWithErrorHandler(logger, errorHandler)
}
