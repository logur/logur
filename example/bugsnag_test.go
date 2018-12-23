package example

import (
	"github.com/bugsnag/bugsnag-go"
	"github.com/goph/logur"
)

func ExampleBugsnag() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	bugsnag.New(bugsnag.Configuration{
		Logger: logur.NewPrintErrorLogger(logger),
	})

	// Output:
}
