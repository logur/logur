package logur

import (
	"context"
)

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
	l.logger.Info(msg, fields...)
}

func (l loggerContext) Warn(msg string, fields ...map[string]interface{}) {
	l.logger.Warn(msg, fields...)
}

func (l loggerContext) Error(msg string, fields ...map[string]interface{}) {
	l.logger.Error(msg, fields...)
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
