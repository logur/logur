/*
Package logrlog provides a github.com/go-logr/logr logger.

With logur you can easily wire the logging library of your choice into any logr logger compatible library:

	package main

	import (
		"github.com/goph/logur"
		"github.com/goph/logur/integrations/logrlog"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		ilogger := logrlog.New(logger)

		// inject the logger somewhere
	}
*/
package logrlog

import (
	logrintegration "logur.dev/integration/logr"

	"github.com/goph/logur"
)

// Logger is a logr Logger.
// Deprecated: use logur.dev/integration/logr.Logger instead.
type Logger = logrintegration.Logger

// New returns a new logr logger.
// Deprecated: use logur.dev/integration/logr.New instead.
func New(logger logur.Logger) *Logger {
	return logrintegration.New(logger)
}
