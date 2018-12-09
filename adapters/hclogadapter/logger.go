package hclogadapter

import (
	"fmt"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/simplelogadapter"
	"github.com/hashicorp/go-hclog"
)

type adapter struct {
	logger hclog.Logger
}

// New returns a new logur compatible logger with hclog as the logging library.
// If nil is passed as logger, the global hclog instance is used as fallback.
func New(logger hclog.Logger) logur.Logger {
	if logger == nil {
		logger = hclog.Default()
	}

	return simplelogadapter.New(&adapter{logger})
}

func (a *adapter) Trace(args ...interface{}) {
	a.logger.Trace(fmt.Sprint(args...))
}

func (a *adapter) Debug(args ...interface{}) {
	a.logger.Debug(fmt.Sprint(args...))
}

func (a *adapter) Info(args ...interface{}) {
	a.logger.Info(fmt.Sprint(args...))
}

func (a *adapter) Warn(args ...interface{}) {
	a.logger.Warn(fmt.Sprint(args...))
}

func (a *adapter) Error(args ...interface{}) {
	a.logger.Error(fmt.Sprint(args...))
}

func (a *adapter) WithFields(fields simplelogadapter.Fields) simplelogadapter.Logger {
	keyvals := make([]interface{}, len(fields)*2)
	i := 0

	for key, value := range fields {
		keyvals[i] = key
		keyvals[i+1] = value

		i += 2
	}

	if keyvals == nil {
		return a
	}

	return &adapter{a.logger.With(keyvals...)}
}
