package zerologadapter

import (
	"fmt"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/simplelogadapter"
	"github.com/rs/zerolog"
)

type adapter struct {
	logger zerolog.Logger
}

// New returns a new logur compatible logger with hclog as the logging library.
// If nil is passed as logger, the global hclog instance is used as fallback.
func New(logger zerolog.Logger) logur.Logger {
	return simplelogadapter.New(&adapter{logger})
}

func (a *adapter) Trace(args ...interface{}) {
	// Fall back to Debug
	a.Debug(args...)
}

func (a *adapter) Debug(args ...interface{}) {
	a.logger.Debug().Msg(fmt.Sprint(args...))
}

func (a *adapter) Info(args ...interface{}) {
	a.logger.Info().Msg(fmt.Sprint(args...))
}

func (a *adapter) Warn(args ...interface{}) {
	a.logger.Warn().Msg(fmt.Sprint(args...))
}

func (a *adapter) Error(args ...interface{}) {
	a.logger.Error().Msg(fmt.Sprint(args...))
}

func (a *adapter) WithFields(fields simplelogadapter.Fields) simplelogadapter.Logger {
	return &adapter{a.logger.With().Fields(fields).Logger()}
}
