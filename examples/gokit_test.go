package example

import (
	"github.com/go-kit/kit/log"

	"github.com/goph/logur"
	"github.com/goph/logur/integrations/kitlog"
)

func Example_goKitLog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	log.With(kitlog.New(logger), "key", "value")

	// Output:
}
