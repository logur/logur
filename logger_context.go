package logur

import (
	"context"
)

type withContextExtractor struct {
	LoggerFacade
	extractor ContextExtractor
}

// WithContextExtractor returns a logger that extracts details from the provided context (if any)
// and annotates the log event with them.
func WithContextExtractor(handler Logger, extractor ContextExtractor) LoggerFacade {
	return withContextExtractor{
		LoggerFacade: ensureLoggerFacade(handler),
		extractor:    extractor,
	}
}

func (l withContextExtractor) TraceContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.LoggerFacade.TraceContext(ctx, msg, mergeFields(l.extractor(ctx), fields))
}

func (l withContextExtractor) DebugContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.LoggerFacade.DebugContext(ctx, msg, mergeFields(l.extractor(ctx), fields))
}

func (l withContextExtractor) InfoContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.LoggerFacade.InfoContext(ctx, msg, mergeFields(l.extractor(ctx), fields))
}

func (l withContextExtractor) WarnContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.LoggerFacade.WarnContext(ctx, msg, mergeFields(l.extractor(ctx), fields))
}

func (l withContextExtractor) ErrorContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.LoggerFacade.ErrorContext(ctx, msg, mergeFields(l.extractor(ctx), fields))
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
