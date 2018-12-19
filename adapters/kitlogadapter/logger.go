// Package kitlogadapter provides a logur compatible adapter for go-kit logger.
package kitlogadapter

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/logur"
	"github.com/goph/logur/internal/keyvals"
)

type adapter struct {
	logger log.Logger
}

// New returns a new logur compatible logger with hclog as the logging library.
// If nil is passed as logger, the global hclog instance is used as fallback.
func New(logger log.Logger) logur.Logger {
	if logger == nil {
		logger = log.NewNopLogger()
	}

	return &adapter{logger}
}

func (a *adapter) Trace(msg string, fields map[string]interface{}) {
	// Fall back to Debug
	a.Debug(msg, fields)
}

func (a *adapter) Debug(msg string, fields map[string]interface{}) {
	_ = level.Debug(a.logger).Log(append(keyvals.FromMap(fields), "msg", msg)...)
}

func (a *adapter) Info(msg string, fields map[string]interface{}) {
	_ = level.Info(a.logger).Log(append(keyvals.FromMap(fields), "msg", msg)...)
}

func (a *adapter) Warn(msg string, fields map[string]interface{}) {
	_ = level.Warn(a.logger).Log(append(keyvals.FromMap(fields), "msg", msg)...)
}

func (a *adapter) Error(msg string, fields map[string]interface{}) {
	_ = level.Error(a.logger).Log(append(keyvals.FromMap(fields), "msg", msg)...)
}
