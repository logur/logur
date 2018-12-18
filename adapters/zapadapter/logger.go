package zapadapter

import (
	"github.com/goph/logur"
	"go.uber.org/zap"
)

type adapter struct {
	*zap.SugaredLogger
}

// New returns a new logur compatible logger with zap as the logging library.
// If nil is passed as logger, the global sugared logger instance is used as fallback.
func New(logger *zap.SugaredLogger) logur.Logger {
	if logger == nil {
		logger = zap.S()
	}

	return &adapter{logger}
}

func (a *adapter) Trace(args ...interface{}) {
	// Fall back to Debug
	a.Debug(args...)
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

	return &adapter{a.SugaredLogger.With(keyvals...)}
}
