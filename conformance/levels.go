package conformance

import (
	"context"

	"logur.dev/logur"
)

// nolint: gochecknoglobals
var testLevelMap = map[logur.Level]struct {
	logFunc    func(logger logur.Logger, msg string, fields ...map[string]interface{})
	logCtxFunc func(logger logur.LoggerContext, ctx context.Context, msg string, fields ...map[string]interface{})
}{
	logur.Trace: {
		logFunc:    logur.Logger.Trace,
		logCtxFunc: logur.LoggerContext.TraceContext,
	},
	logur.Debug: {
		logFunc:    logur.Logger.Debug,
		logCtxFunc: logur.LoggerContext.DebugContext,
	},
	logur.Info: {
		logFunc:    logur.Logger.Info,
		logCtxFunc: logur.LoggerContext.InfoContext,
	},
	logur.Warn: {
		logFunc:    logur.Logger.Warn,
		logCtxFunc: logur.LoggerContext.WarnContext,
	},
	logur.Error: {
		logFunc:    logur.Logger.Error,
		logCtxFunc: logur.LoggerContext.ErrorContext,
	},
}

// nolint: gochecknoglobals
var allLevels = []logur.Level{logur.Trace, logur.Debug, logur.Info, logur.Warn, logur.Error}
