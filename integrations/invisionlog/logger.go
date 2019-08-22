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
	invisionintegration "logur.dev/integration/invision"

	"github.com/goph/logur"
)

// Logger is a github.com/InVisionApp/go-logger.Logger logger.
// Deprecated: use logur.dev/integration/invision.Logger instead.
type Logger = invisionintegration.Logger

// New returns a new github.com/InVisionApp/go-logger.Logger logger.
// Deprecated: use logur.dev/integration/invision.New instead.
func New(logger logur.Logger) *Logger {
	return invisionintegration.New(logger)
}
