package example

import (
	"github.com/bugsnag/bugsnag-go"

	"logur.dev/logur"
)

func Example_bugsnag() {
	logger := logur.NoopLogger{} // choose an actual implementation

	bugsnag.New(bugsnag.Configuration{
		Logger: logur.NewErrorPrintLogger(logger),
	})

	// Output:
}
