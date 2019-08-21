// Package hclogadapter provides a logur adapter for hclog.
package hclogadapter

import (
	"github.com/hashicorp/go-hclog"
	hclogadapter "logur.dev/adapter/hclog"
)

// Logger is a logur adapter for hclog.
// Deprecated: use logur.dev/adapter/hclog.Logger instead.
type Logger = hclogadapter.Logger

// New returns a new logur compatible logger with hclog as the logging library.
// If nil is passed as logger, the global hclog instance is used as fallback.
// Deprecated: use logur.dev/adapter/hclog.New instead.
func New(logger hclog.Logger) *Logger {
	return hclogadapter.New(logger)
}
