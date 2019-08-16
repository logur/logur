package example

import (
	"github.com/go-kit/kit/log"

	"github.com/goph/logur"
)

func Example_goKitLog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	log.With(logur.NewKitLogger(logger), "key", "value")

	// Output:
}
