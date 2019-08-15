package example

import (
	"github.com/rollbar/rollbar-go"

	"github.com/goph/logur"
)

func Example_rollbar() {
	logger := logur.NewNoopLogger() // choose an actual implementation
	clientLogger := logur.NewErrorPrintLogger(logger)

	rollbar.SetLogger(clientLogger)
	// OR
	notifier := rollbar.New("token", "environment", "version", "host", "root")
	notifier.SetLogger(clientLogger)

	// Output:
}
