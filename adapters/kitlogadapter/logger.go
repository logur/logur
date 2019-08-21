// Package kitlogadapter provides a logur adapter for go-kit logger.
package kitlogadapter

import (
	"github.com/go-kit/kit/log"
	kitadapter "logur.dev/adapter/kit"
)

// Logger is a logur adapter for go-kit logger.
// Deprecated: use logur.dev/adapter/kit.Logger instead.
type Logger = kitadapter.Logger

// New returns a new logur compatible logger with go-kit log as the logging library.
// If nil is passed as logger, a noop logger is used as fallback.
// Deprecated: use logur.dev/adapter/kit.New instead.
func New(logger log.Logger) *Logger {
	return kitadapter.New(logger)
}
