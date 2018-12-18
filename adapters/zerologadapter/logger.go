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

func (a *adapter) Trace(msg string) {
	// Fall back to Debug
	a.Debug(msg)
}

func (a *adapter) Debug(msg string) {
	a.logger.Debug().Msg(msg)
}

func (a *adapter) Info(msg string) {
	a.logger.Info().Msg(msg)
}

func (a *adapter) Warn(msg string) {
	a.logger.Warn().Msg(msg)
}

func (a *adapter) Error(msg string) {
	a.logger.Error().Msg(msg)
}

func (a *adapter) WithFields(fields map[string]interface{}) logur.Logger {
	return &adapter{a.logger.With().Fields(fields).Logger()}
}
