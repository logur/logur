// Package zapadapter provides a logur adapter for Uber's Zap.
package zapadapter

import (
	_ "go.uber.org/multierr" // Required because multierr is not a go module.
	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"
)

// Logger is a logur adapter for Uber's zap.
// Deprecated: use logur.dev/adapter/zap.Logger instead.
type Logger = zapadapter.Logger

// New returns a new logur compatible logger with Uber's zap as the logging library.
// If nil is passed as logger, the global sugared logger instance is used as fallback.
// Deprecated: use logur.dev/adapter/zap.New instead.
func New(logger *zap.Logger) *Logger {
	return zapadapter.New(logger)
}
