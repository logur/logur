// Package zerologadapter provides a logur adapter for zerolog.
package zerologadapter

import (
	"github.com/rs/zerolog"
	zerologadapter "logur.dev/adapter/zerolog"
)

// Logger is a logur adapter for zerolog.
// Deprecated: use logur.dev/adapter/zerolog.Logger instead.
type Logger = zerologadapter.Logger

// New returns a new logur compatible logger with zerolog as the logging library.
// Deprecated: use logur.dev/adapter/zerolog.New instead.
func New(logger zerolog.Logger) *Logger {
	return zerologadapter.New(logger)
}
