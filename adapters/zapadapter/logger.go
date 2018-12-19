// Package zapadapter provides a logur compatible adapter for Uber's Zap.
package zapadapter

import (
	"github.com/goph/logur"
	"github.com/goph/logur/internal/keyvals"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type adapter struct {
	logger *zap.SugaredLogger
	core   zapcore.Core
}

// New returns a new logur compatible logger with zap as the logging library.
// If nil is passed as logger, the global sugared logger instance is used as fallback.
func New(logger *zap.Logger) logur.Logger {
	if logger == nil {
		logger = zap.L()
	}

	return &adapter{
		logger: logger.Sugar(),
		core:   logger.Core(),
	}
}

func (a *adapter) Trace(msg string, fields map[string]interface{}) {
	// Fall back to Debug
	a.Debug(msg, fields)
}

func (a *adapter) Debug(msg string, fields map[string]interface{}) {
	if !a.core.Enabled(zap.DebugLevel) {
		return
	}

	a.logger.Debugw(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Info(msg string, fields map[string]interface{}) {
	if !a.core.Enabled(zap.InfoLevel) {
		return
	}

	a.logger.Infow(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Warn(msg string, fields map[string]interface{}) {
	if !a.core.Enabled(zap.WarnLevel) {
		return
	}

	a.logger.Warnw(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Error(msg string, fields map[string]interface{}) {
	if !a.core.Enabled(zap.ErrorLevel) {
		return
	}

	a.logger.Errorw(msg, keyvals.FromMap(fields)...)
}

// nolint: gochecknoglobals
var levelMap = map[logur.Level]zapcore.Level{
	logur.Trace: zap.DebugLevel,
	logur.Debug: zap.DebugLevel,
	logur.Info:  zap.InfoLevel,
	logur.Warn:  zap.WarnLevel,
	logur.Error: zap.ErrorLevel,
}

func (a *adapter) LevelEnabled(level logur.Level) bool {
	checkLevel, ok := levelMap[level]
	if !ok {
		return true
	}

	return a.core.Enabled(checkLevel)
}
