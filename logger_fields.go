package logur

import (
	"context"
)

// WithFields returns a new logger instance that attaches the given fields to every subsequent log call.
func WithFields(logger Logger, fields map[string]interface{}) LoggerFacade {
	loggerFacade := ensureLoggerFacade(logger)

	if len(fields) == 0 {
		return loggerFacade
	}

	// Do not add a new layer
	// Create a new logger instead with the parent fields
	//
	// fieldLogger already implements LoggerFacade, so loggerFacade should be the same as logger if it's a fieldLogger
	if l, ok := loggerFacade.(*fieldLogger); ok && len(l.fields) > 0 {
		_fields := make(map[string]interface{}, len(l.fields)+len(fields))

		for key, value := range l.fields {
			_fields[key] = value
		}

		for key, value := range fields {
			_fields[key] = value
		}

		fields = _fields
		logger = l.logger
	}

	l := &fieldLogger{logger: loggerFacade, fields: fields}

	if levelEnabler, ok := logger.(LevelEnabler); ok {
		l.levelEnabler = levelEnabler
	}

	return l
}

// WithField is a shortcut for WithFields(logger, map[string]interface{}{key: value}).
func WithField(logger Logger, key string, value interface{}) Logger {
	return WithFields(logger, map[string]interface{}{key: value})
}

// fieldLogger holds a context and passes it to the underlying logger when a log event is recorded.
type fieldLogger struct {
	logger       LoggerFacade
	fields       map[string]interface{}
	levelEnabler LevelEnabler
}

// Trace implements the logur.Logger interface.
func (l *fieldLogger) Trace(msg string, fields ...map[string]interface{}) {
	l.log(Trace, l.logger.Trace, msg, fields)
}

// Debug implements the logur.Logger interface.
func (l *fieldLogger) Debug(msg string, fields ...map[string]interface{}) {
	l.log(Debug, l.logger.Debug, msg, fields)
}

// Info implements the logur.Logger interface.
func (l *fieldLogger) Info(msg string, fields ...map[string]interface{}) {
	l.log(Info, l.logger.Info, msg, fields)
}

// Warn implements the logur.Logger interface.
func (l *fieldLogger) Warn(msg string, fields ...map[string]interface{}) {
	l.log(Warn, l.logger.Warn, msg, fields)
}

// Error implements the logur.Logger interface.
func (l *fieldLogger) Error(msg string, fields ...map[string]interface{}) {
	l.log(Error, l.logger.Error, msg, fields)
}

// log deduplicates some field logger code.
func (l *fieldLogger) log(level Level, fn LogFunc, msg string, fields []map[string]interface{}) {
	if !l.levelEnabled(level) {
		return
	}

	fn(msg, mergeFields(l.fields, fields))
}

func (l *fieldLogger) TraceContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logContext(Trace, l.logger.TraceContext, ctx, msg, fields)
}

func (l *fieldLogger) DebugContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logContext(Debug, l.logger.DebugContext, ctx, msg, fields)
}

func (l *fieldLogger) InfoContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logContext(Info, l.logger.InfoContext, ctx, msg, fields)
}

func (l *fieldLogger) WarnContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logContext(Warn, l.logger.WarnContext, ctx, msg, fields)
}

func (l *fieldLogger) ErrorContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logContext(Error, l.logger.ErrorContext, ctx, msg, fields)
}

// logContext deduplicates some field logger code.
// nolint: golint
func (l *fieldLogger) logContext(
	level Level,
	fn LogContextFunc,
	ctx context.Context,
	msg string,
	fields []map[string]interface{},
) {
	if !l.levelEnabled(level) {
		return
	}

	fn(ctx, msg, mergeFields(l.fields, fields))
}

func (l *fieldLogger) levelEnabled(level Level) bool {
	if l.levelEnabler != nil {
		return l.levelEnabler.LevelEnabled(level)
	}

	return true
}
