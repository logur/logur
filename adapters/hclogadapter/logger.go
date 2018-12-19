// Package hclogadapter provides a logur compatible adapter for hclog.
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
	if !a.logger.IsTrace() {
		return
	}

	a.logger.Trace(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Debug(msg string, fields map[string]interface{}) {
	if !a.logger.IsDebug() {
		return
	}

	a.logger.Debug(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Info(msg string, fields map[string]interface{}) {
	if !a.logger.IsInfo() {
		return
	}

	a.logger.Info(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Warn(msg string, fields map[string]interface{}) {
	if !a.logger.IsWarn() {
		return
	}

	a.logger.Warn(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Error(msg string, fields map[string]interface{}) {
	if !a.logger.IsError() {
		return
	}

	a.logger.Error(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) LevelEnabled(level logur.Level) bool {
	switch level {
	case logur.Trace:
		return a.logger.IsTrace()
	case logur.Debug:
		return a.logger.IsDebug()
	case logur.Info:
		return a.logger.IsInfo()
	case logur.Warn:
		return a.logger.IsWarn()
	case logur.Error:
		return a.logger.IsError()
	}

	return true
}
