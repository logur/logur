package example

import (
	"google.golang.org/grpc/grpclog"

	"logur.dev/logur"
	grpcintegration "logur.dev/logur/integration/grpc"
)

func Example_grpcLog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	grpclog.SetLoggerV2(grpcintegration.New(logger))

	// Output:
}
