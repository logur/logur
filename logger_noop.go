package logur

type noopLogger struct{}

// NewNoopLogger creates a no-op logger that discards all received log events.
// Useful in examples and as a fallback logger.
func NewNoopLogger() Logger {
	return &noopLogger{}
}

func (*noopLogger) Trace(msg string, fields map[string]interface{}) {}
func (*noopLogger) Debug(msg string, fields map[string]interface{}) {}
func (*noopLogger) Info(msg string, fields map[string]interface{})  {}
func (*noopLogger) Warn(msg string, fields map[string]interface{})  {}
func (*noopLogger) Error(msg string, fields map[string]interface{}) {}
