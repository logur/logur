package logur

import (
	"context"
	"fmt"
	"strings"
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
// for returned handlers. It's not supposed to be used by consumers of this package.
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

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	fn(msg, f)
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

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	fn(ctx, msg, f)
}

func (l *fieldLogger) mergeFields(fields map[string]interface{}) map[string]interface{} {
	if len(fields) == 0 { // Not having any fields passed to the log function has a higher chance
		return l.fields
	}

	if len(l.fields) == 0 { // This is possible too, but has a much lower probability
		return fields
	}

	f := make(map[string]interface{}, len(fields)+len(l.fields))

	for key, value := range l.fields {
		f[key] = value
	}

	for key, value := range fields {
		f[key] = value
	}

	return f
}

func (l *fieldLogger) levelEnabled(level Level) bool {
	if l.levelEnabler != nil {
		return l.levelEnabler.LevelEnabled(level)
	}

	return true
}

type loggerContext struct {
	logger    LoggerFacade
	extractor ContextExtractor
}

// NewLoggerContext returns an logger that extracts details from the provided context (if any)
// and annotates the log event with them.
func NewLoggerContext(handler Logger, extractor ContextExtractor) LoggerFacade {
	return loggerContext{
		logger:    ensureLoggerFacade(handler),
		extractor: extractor,
	}
}

func (l loggerContext) Trace(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}

func (l loggerContext) Debug(msg string, fields ...map[string]interface{}) {
	l.logger.Debug(msg, fields...)
}

func (l loggerContext) Info(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}

func (l loggerContext) Warn(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}

func (l loggerContext) Error(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}

func (l loggerContext) TraceContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logger.TraceContext(ctx, msg, l.mergeFields(l.extractor(ctx), fields))
}

func (l loggerContext) DebugContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logger.DebugContext(ctx, msg, l.mergeFields(l.extractor(ctx), fields))
}

func (l loggerContext) InfoContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logger.InfoContext(ctx, msg, l.mergeFields(l.extractor(ctx), fields))
}

func (l loggerContext) WarnContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logger.WarnContext(ctx, msg, l.mergeFields(l.extractor(ctx), fields))
}

func (l loggerContext) ErrorContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.logger.ErrorContext(ctx, msg, l.mergeFields(l.extractor(ctx), fields))
}

func (l loggerContext) mergeFields(ctxFields Fields, fields []map[string]interface{}) Fields {
	if len(fields) == 0 {
		return ctxFields
	}

	if len(ctxFields) == 0 {
		return fields[0]
	}

	// the maximum length of the map is the sum of the two map's length
	f := make(map[string]interface{}, len(fields)+len(fields[0]))

	for key, value := range ctxFields {
		f[key] = value
	}

	for key, value := range fields[0] {
		f[key] = value
	}

	return f
}

// ContextExtractor extracts a map of details from a context.
type ContextExtractor func(ctx context.Context) map[string]interface{}

// ContextExtractors combines a list of ContextExtractor.
// The returned extractor aggregates the result of the underlying extractors.
func ContextExtractors(extractors ...ContextExtractor) ContextExtractor {
	return func(ctx context.Context) map[string]interface{} {
		fields := make(map[string]interface{})

		for _, extractor := range extractors {
			for key, value := range extractor(ctx) {
				fields[key] = value
			}
		}

		return fields
	}
}

// PrintLogger logs messages with fmt.Print* function semantics.
type PrintLogger struct {
	logger LogFunc
}

// NewPrintLogger returns a new PrintLogger.
func NewPrintLogger(logger LogFunc) *PrintLogger {
	return &PrintLogger{logger}
}

// NewErrorPrintLogger returns a new PrintLogger that logs everything on error level.
func NewErrorPrintLogger(logger Logger) *PrintLogger {
	return NewPrintLogger(LevelFunc(logger, Error))
}

// Print logs a line with fmt.Print semantics.
func (l *PrintLogger) Print(v ...interface{}) {
	l.logger(fmt.Sprint(v...))
}

// Println logs a line with fmt.Println semantics.
func (l *PrintLogger) Println(v ...interface{}) {
	l.logger(strings.TrimSuffix(fmt.Sprintln(v...), "\n"))
}

// Printf logs a line with fmt.Printf semantics.
func (l *PrintLogger) Printf(format string, args ...interface{}) {
	l.logger(fmt.Sprintf(format, args...))
}

// MessageLogger simplifies the Logger interface by removing the second context parameter.
// Useful when there is no need for contextual logging.
type MessageLogger struct {
	logger Logger
}

// NewMessageLogger returns a new MessageLogger instance.
func NewMessageLogger(logger Logger) *MessageLogger {
	return &MessageLogger{logger}
}

// Trace logs a Trace level event.
func (l *MessageLogger) Trace(msg string) { l.logger.Trace(msg) }

// Debug logs a Debug level event.
func (l *MessageLogger) Debug(msg string) { l.logger.Debug(msg) }

// Info logs a Info level event.
func (l *MessageLogger) Info(msg string) { l.logger.Info(msg) }

// Warn logs a Warn level event.
func (l *MessageLogger) Warn(msg string) { l.logger.Warn(msg) }

// Error logs a Error level event.
func (l *MessageLogger) Error(msg string) { l.logger.Error(msg) }
