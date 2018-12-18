package kitlogadapter

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/logur"
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

func (a *adapter) Trace(args ...interface{}) {
	// Fall back to Debug
	a.Debug(args...)
}

func (a *adapter) Debug(args ...interface{}) {
	_ = level.Debug(a.logger).Log("msg", fmt.Sprint(args...))
}

func (a *adapter) Info(args ...interface{}) {
	_ = level.Info(a.logger).Log("msg", fmt.Sprint(args...))
}

func (a *adapter) Warn(args ...interface{}) {
	_ = level.Warn(a.logger).Log("msg", fmt.Sprint(args...))
}

func (a *adapter) Error(args ...interface{}) {
	_ = level.Error(a.logger).Log("msg", fmt.Sprint(args...))
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

	return &adapter{log.With(a.logger, keyvals...)}
}
