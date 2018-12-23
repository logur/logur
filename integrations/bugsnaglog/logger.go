/*
Package bugsnaglog provides a logger for Bugsnag's error notifier.

With logur you can easily wire the logging library of your choice into Bugsnag's notifier:

	package main

	import (
		"github.com/goph/logur"
		"github.com/goph/logur/integrations/bugsnaglog"
		"github.com/bugsnag/bugsnag-go"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		config := bugsnag.Configuration{
			Logger: bugsnaglog.New(logger),
		}
	}
*/
package bugsnaglog

import (
	"github.com/goph/logur"
)

// New returns a new logger for Bugsnag's error notifier.
func New(logger logur.Logger) *logur.PrintLogger {
	return logur.NewPrintErrorLogger(logger)
}
