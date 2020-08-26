package logur

import (
	"context"

	kvs "logur.dev/logur/internal/keyvals"
)

// Logger is a unified interface for various logging use cases and practices, including:
// 		- leveled logging
// 		- structured logging
type Logger interface {
	// Trace logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace(msg string, fields ...map[string]interface{})

	// Debug logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug(msg string, fields ...map[string]interface{})

	// Info logs an Info event.
	//
	// General information about what's happening inside the system.
	Info(msg string, fields ...map[string]interface{})

	// Warn logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	Warn(msg string, fields ...map[string]interface{})

	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	Error(msg string, fields ...map[string]interface{})
}

// LoggerContext is an optional interface that MAY be implemented by a Logger.
// It is similar to Logger, but it receives a context as the first parameter.
// An implementation MAY extract information from the context and annotate the log context with it.
//
// LoggerContext MAY honor the deadline carried by the context, but that's not a hard requirement.
type LoggerContext interface {
	// TraceContext logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	TraceContext(ctx context.Context, msg string, fields ...map[string]interface{})

	// DebugContext logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	DebugContext(ctx context.Context, msg string, fields ...map[string]interface{})

	// InfoContext logs an Info event.
	//
	// General information about what's happening inside the system.
	InfoContext(ctx context.Context, msg string, fields ...map[string]interface{})

	// WarnContext logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	WarnContext(ctx context.Context, msg string, fields ...map[string]interface{})

	// ErrorContext logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	ErrorContext(ctx context.Context, msg string, fields ...map[string]interface{})
}

// LoggerFacade is a combination of Logger and LoggerContext.
// It's sole purpose is to make the API of the package concise by exposing a common interface type
// for returned loggers. It's not supposed to be used by consumers of this package.
//
// It goes directly against the "Use interfaces, return structs" idiom of Go,
// but at the current phase of the package the smaller API surface makes more sense.
//
// In the future it might get replaced with concrete types.
type LoggerFacade interface {
	Logger
	LoggerContext
}

func ensureLoggerFacade(logger Logger) LoggerFacade {
	if logger, ok := logger.(LoggerFacade); ok {
		return logger
	}

	if levelEnabler, ok := logger.(LevelEnabler); ok {
		return levelEnablerLoggerFacade{
			LoggerFacade: loggerFacade{logger},
			LevelEnabler: levelEnabler,
		}
	}

	return loggerFacade{logger}
}

type loggerFacade struct {
	Logger
}

func (l loggerFacade) TraceContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Trace(msg, fields...)
}

func (l loggerFacade) DebugContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Debug(msg, fields...)
}

func (l loggerFacade) InfoContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Info(msg, fields...)
}

func (l loggerFacade) WarnContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Warn(msg, fields...)
}

func (l loggerFacade) ErrorContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Error(msg, fields...)
}

type levelEnablerLoggerFacade struct {
	LoggerFacade
	LevelEnabler
}

// Fields is used to define structured fields which are appended to log events.
// It can be used as a shorthand for map[string]interface{}.
type Fields map[string]interface{}

// LogFunc records a log event.
type LogFunc func(msg string, fields ...map[string]interface{})

// LogContextFunc records a log event.
type LogContextFunc func(ctx context.Context, msg string, fields ...map[string]interface{})

// NoopLogger is a no-op logger that discards all received log events.
//
// It implements both Logger and LoggerContext interfaces.
type NoopLogger struct{}

func (NoopLogger) Trace(_ string, _ ...map[string]interface{}) {}
func (NoopLogger) Debug(_ string, _ ...map[string]interface{}) {}
func (NoopLogger) Info(_ string, _ ...map[string]interface{})  {}
func (NoopLogger) Warn(_ string, _ ...map[string]interface{})  {}
func (NoopLogger) Error(_ string, _ ...map[string]interface{}) {}

func (NoopLogger) TraceContext(_ context.Context, _ string, _ ...map[string]interface{}) {}
func (NoopLogger) DebugContext(_ context.Context, _ string, _ ...map[string]interface{}) {}
func (NoopLogger) InfoContext(_ context.Context, _ string, _ ...map[string]interface{})  {}
func (NoopLogger) WarnContext(_ context.Context, _ string, _ ...map[string]interface{})  {}
func (NoopLogger) ErrorContext(_ context.Context, _ string, _ ...map[string]interface{}) {}

// NewNoopLogger creates a no-op logger that discards all received log events.
// Useful in examples and as a fallback logger.
//
// Deprecated: use NoopLogger.
func NewNoopLogger() Logger {
	return NoopLogger{}
}

// KVLogger is a unified interface for various logging use cases and practices, including:
// 		- leveled logging
// 		- structured logging
//
// Compared to Logger, KVLogger accepts key-value pairs in the form of variadic interface arguments.
type KVLogger interface {
	// Trace logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace(msg string, keyvals ...interface{})

	// Debug logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug(msg string, keyvals ...interface{})

	// Info logs an Info event.
	//
	// General information about what's happening inside the system.
	Info(msg string, keyvals ...interface{})

	// Warn logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	Warn(msg string, keyvals ...interface{})

	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	Error(msg string, keyvals ...interface{})
}

// KVLoggerContext is an optional interface that MAY be implemented by a KVLogger.
// It is similar to KVLogger, but it receives a context as the first parameter.
// An implementation MAY extract information from the context and annotate the log context with it.
//
// KVLoggerContext MAY honor the deadline carried by the context, but that's not a hard requirement.
type KVLoggerContext interface {
	// TraceContext logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	TraceContext(ctx context.Context, msg string, keyvals ...interface{})

	// DebugContext logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	DebugContext(ctx context.Context, msg string, keyvals ...interface{})

	// InfoContext logs an Info event.
	//
	// General information about what's happening inside the system.
	InfoContext(ctx context.Context, msg string, keyvals ...interface{})

	// WarnContext logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	WarnContext(ctx context.Context, msg string, keyvals ...interface{})

	// ErrorContext logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	ErrorContext(ctx context.Context, msg string, keyvals ...interface{})
}

// KVLoggerFacade is a combination of KVLogger and KVLoggerContext.
// It's sole purpose is to make the API of the package concise by exposing a common interface type
// for returned loggers. It's not supposed to be used by consumers of this package.
//
// It goes directly against the "Use interfaces, return structs" idiom of Go,
// but at the current phase of the package the smaller API surface makes more sense.
//
// In the future it might get replaced with concrete types.
type KVLoggerFacade interface {
	KVLogger
	KVLoggerContext
}

// LoggerToKV converts a Logger to a KVLogger.
func LoggerToKV(logger Logger) KVLoggerFacade {
	return loggerToKV{logger: ensureLoggerFacade(logger)}
}

type loggerToKV struct {
	logger LoggerFacade
}

func (l loggerToKV) Trace(msg string, keyvals ...interface{}) {
	l.logger.Trace(msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) Debug(msg string, keyvals ...interface{}) {
	l.logger.Debug(msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) Info(msg string, keyvals ...interface{}) {
	l.logger.Info(msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) Warn(msg string, keyvals ...interface{}) {
	l.logger.Warn(msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) Error(msg string, keyvals ...interface{}) {
	l.logger.Error(msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) TraceContext(ctx context.Context, msg string, keyvals ...interface{}) {
	l.logger.TraceContext(ctx, msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) DebugContext(ctx context.Context, msg string, keyvals ...interface{}) {
	l.logger.DebugContext(ctx, msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) InfoContext(ctx context.Context, msg string, keyvals ...interface{}) {
	l.logger.InfoContext(ctx, msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) WarnContext(ctx context.Context, msg string, keyvals ...interface{}) {
	l.logger.WarnContext(ctx, msg, kvs.ToMap(keyvals))
}

func (l loggerToKV) ErrorContext(ctx context.Context, msg string, keyvals ...interface{}) {
	l.logger.ErrorContext(ctx, msg, kvs.ToMap(keyvals))
}

func ensureKVLoggerFacade(logger KVLogger) KVLoggerFacade {
	if logger, ok := logger.(KVLoggerFacade); ok {
		return logger
	}

	if levelEnabler, ok := logger.(LevelEnabler); ok {
		return levelEnablerKVLoggerFacade{
			KVLoggerFacade: kvLoggerFacade{logger},
			LevelEnabler:   levelEnabler,
		}
	}

	return kvLoggerFacade{logger}
}

type kvLoggerFacade struct {
	KVLogger
}

func (l kvLoggerFacade) TraceContext(_ context.Context, msg string, keyvals ...interface{}) {
	l.Trace(msg, keyvals...)
}

func (l kvLoggerFacade) DebugContext(_ context.Context, msg string, keyvals ...interface{}) {
	l.Debug(msg, keyvals...)
}

func (l kvLoggerFacade) InfoContext(_ context.Context, msg string, keyvals ...interface{}) {
	l.Info(msg, keyvals...)
}

func (l kvLoggerFacade) WarnContext(_ context.Context, msg string, keyvals ...interface{}) {
	l.Warn(msg, keyvals...)
}

func (l kvLoggerFacade) ErrorContext(_ context.Context, msg string, keyvals ...interface{}) {
	l.Error(msg, keyvals...)
}

type levelEnablerKVLoggerFacade struct {
	KVLoggerFacade
	LevelEnabler
}

// KVLogFunc records a log event.
type KVLogFunc func(msg string, keyvals ...interface{})

// KVLogContextFunc records a log event.
type KVLogContextFunc func(ctx context.Context, msg string, keyvals ...interface{})

// NoopKVLogger is a no-op logger that discards all received log events.
//
// It implements both KVLogger and KVLoggerContext interfaces.
type NoopKVLogger struct{}

func (NoopKVLogger) Trace(_ string, _ ...interface{}) {}
func (NoopKVLogger) Debug(_ string, _ ...interface{}) {}
func (NoopKVLogger) Info(_ string, _ ...interface{})  {}
func (NoopKVLogger) Warn(_ string, _ ...interface{})  {}
func (NoopKVLogger) Error(_ string, _ ...interface{}) {}

func (NoopKVLogger) TraceContext(_ context.Context, _ string, _ ...interface{}) {}
func (NoopKVLogger) DebugContext(_ context.Context, _ string, _ ...interface{}) {}
func (NoopKVLogger) InfoContext(_ context.Context, _ string, _ ...interface{})  {}
func (NoopKVLogger) WarnContext(_ context.Context, _ string, _ ...interface{})  {}
func (NoopKVLogger) ErrorContext(_ context.Context, _ string, _ ...interface{}) {}
