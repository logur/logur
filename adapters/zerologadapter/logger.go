// Package zerologadapter provides a logur compatible adapter for zerolog.
package zerologadapter

import (
	"github.com/goph/logur"
	"github.com/rs/zerolog"
)

type adapter struct {
	logger zerolog.Logger
}

// New returns a new logur compatible logger with hclog as the logging library.
// If nil is passed as logger, the global hclog instance is used as fallback.
func New(logger zerolog.Logger) logur.Logger {
	return &adapter{logger}
}

func (a *adapter) Trace(msg string, fields map[string]interface{}) {
	// Fall back to Debug
	a.Debug(msg, fields)
}

func (a *adapter) Debug(msg string, fields map[string]interface{}) {
	a.logger.Debug().Fields(fields).Msg(msg)
}

func (a *adapter) Info(msg string, fields map[string]interface{}) {
	a.logger.Info().Fields(fields).Msg(msg)
}

func (a *adapter) Warn(msg string, fields map[string]interface{}) {
	a.logger.Warn().Fields(fields).Msg(msg)
}

func (a *adapter) Error(msg string, fields map[string]interface{}) {
	a.logger.Error().Fields(fields).Msg(msg)
}
