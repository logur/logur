package hclogadapter

import (
	"github.com/goph/logur"
	"github.com/goph/logur/internal/keyvals"
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

	return &adapter{logger}
}

func (a *adapter) Trace(msg string, fields map[string]interface{}) {
	a.logger.Trace(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Debug(msg string, fields map[string]interface{}) {
	a.logger.Debug(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Info(msg string, fields map[string]interface{}) {
	a.logger.Info(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Warn(msg string, fields map[string]interface{}) {
	a.logger.Warn(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Error(msg string, fields map[string]interface{}) {
	a.logger.Error(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) WithFields(fields map[string]interface{}) logur.Logger {
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
