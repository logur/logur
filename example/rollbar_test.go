package example

import (
	"github.com/goph/logur"
	rollbar "github.com/rollbar/rollbar-go"
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
