/*
Package kitlog provides a go-kit logger.

With logur you can easily wire the logging library of your choice into go-kit:

	package main

	import (
		"github.com/goph/logur"
		"github.com/goph/logur/integrations/kitlog"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		kitlogger := kitlog.New(logger)

		// inject the logger somewhere
	}
*/
package kitlog

import (
	"github.com/goph/logur"
)

// New returns a new go-kit logger.
// Deprecated: use github.com/goph/logur.NewKitLogger instead.
func New(logger logur.Logger) *logur.KitLogger {
	return logur.NewKitLogger(logger)
}
