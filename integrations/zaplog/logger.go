package zaplog

import (
	"go.uber.org/zap"
	zapintegration "logur.dev/integration/zap"

	"github.com/goph/logur"
)

// New returns a new zap logger.
// Deprecated: use logur.dev/integration/zap.New instead.
func New(logger logur.Logger) *zap.Logger {
	return zapintegration.New(logger)
}
