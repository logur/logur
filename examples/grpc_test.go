package example

import (
	"google.golang.org/grpc/grpclog"

	"github.com/goph/logur"
	"github.com/goph/logur/integrations/grpcloglog"
)

func Example_grpcLog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	grpclog.SetLoggerV2(grpcloglog.New(logger))

	// Output:
}
