package logur

type noopLogger struct{}

// NewNoopLogger creates a no-op logger that can be used to silence
// all logging from this library. Also useful in tests.
func NewNoopLogger() Logger {
	return &noopLogger{}
}

func (*noopLogger) Trace(msg string, fields map[string]interface{}) {}
func (*noopLogger) Debug(msg string, fields map[string]interface{}) {}
func (*noopLogger) Info(msg string, fields map[string]interface{})  {}
func (*noopLogger) Warn(msg string, fields map[string]interface{})  {}
func (*noopLogger) Error(msg string, fields map[string]interface{}) {}
