package grpcloglog_test

import (
	"google.golang.org/grpc/grpclog"

	"github.com/goph/logur"
	"github.com/goph/logur/integrations/grpcloglog"
)

func Example_grpclog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	grpclog.SetLoggerV2(grpcloglog.New(logger))

	// Output:
}
