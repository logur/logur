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
		logger := logur.NewNoop() // choose an actual implementation
		config := bugsnag.Configuration{
			Logger: bugsnaglog.New(logger),
		}
	}
*/
package bugsnaglog

import (
	"fmt"

	"github.com/goph/logur"
)

// Logger is a logger for Bugsnag's error notifier.
type Logger struct {
	logger logur.Logger
}

// New returns a new logger for Bugsnag's error notifier.
func New(logger logur.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Printf logs bugsnag's errors.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, v...), nil)
}
