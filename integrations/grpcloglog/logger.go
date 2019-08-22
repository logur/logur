/*
Package grpcloglog provides a V2 gRPC logger.

gRPC operates with a globally configured logger that implements the google.golang.org/grpc/grpclog.LoggerV2 interface.

With logur you can easily wire the logging library of your choice into gRPC:

	package main

	import (
		"github.com/goph/logur"
		"github.com/goph/logur/integrations/grpcloglog"
		"google.golang.org/grpc/grpclog"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		grpclog.SetLoggerV2(grpcloglog.New(logger))
	}
*/
package grpcloglog

import (
	"github.com/goph/logur"
	grpcintegration "github.com/goph/logur/integration/grpc"
)

// Logger is a V2 gRPC logger.
type Logger = grpcintegration.Logger

// New returns a new V2 gRPC logger.
func New(logger logur.Logger) *Logger {
	return grpcintegration.New(logger)
}
