package grpcloglog_test

import (
	"github.com/goph/logur"
	"github.com/goph/logur/integrations/grpcloglog"
	"google.golang.org/grpc/grpclog"
)

func Example_grpclog() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	grpclog.SetLoggerV2(grpcloglog.New(logger))

	// Output:
}
