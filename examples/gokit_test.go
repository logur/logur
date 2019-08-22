package example

import (
	"github.com/go-kit/kit/log"

	"github.com/goph/logur"
	kitintegration "github.com/goph/logur/integration/kit"
)

func Example_goKitLog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	log.With(kitintegration.New(logger), "key", "value")

	// Output:
}
