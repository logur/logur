package example

import (
	"google.golang.org/grpc/grpclog"

	"github.com/goph/logur"
)

func Example_grpcLog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	grpclog.SetLoggerV2(logur.NewGRPCV2Logger(logger))

	// Output:
}
