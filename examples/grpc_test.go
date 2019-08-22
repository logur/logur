package example

import (
	"google.golang.org/grpc/grpclog"

	"github.com/goph/logur"
	grpcintegration "github.com/goph/logur/integration/grpc"
)

func Example_grpcLog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	grpclog.SetLoggerV2(grpcintegration.New(logger))

	// Output:
}
