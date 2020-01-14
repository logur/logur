package example

import (
	"github.com/go-kit/kit/log"

	"logur.dev/logur"
	kitintegration "logur.dev/logur/integration/kit"
)

func Example_goKitLog() {
	logger := logur.NoopLogger{} // choose an actual implementation

	log.With(kitintegration.New(logger), "key", "value")

	// Output:
}
