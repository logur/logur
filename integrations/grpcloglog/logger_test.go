package grpcloglog

import (
	"testing"

	"github.com/goph/logur"
	"google.golang.org/grpc/grpclog"
)

func TestLogger(t *testing.T) {
	var _ grpclog.LoggerV2 = New(logur.NewNoopLogger())
}
