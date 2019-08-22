/*
Package kitlog provides a go-kit logger.

With logur you can easily wire the logging library of your choice into go-kit:

	package main

	import (
		"github.com/goph/logur"
		kitintegration "github.com/goph/logur/integration/kit"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		kitlogger := kitintegration.New(logger)

		// inject the logger somewhere
	}
*/
package kitlog

import (
	"github.com/goph/logur"
	kitintegration "github.com/goph/logur/integration/kit"
)

// Logger is a go-kit logger.
type Logger = kitintegration.Logger

// New returns a new go-kit logger.
func New(logger logur.Logger) *Logger {
	return kitintegration.New(logger)
}
