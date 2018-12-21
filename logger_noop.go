package logur

type NoopLogger struct{}

// NewNoop creates a no-op logger that can be used to silence
// all logging from this library. Also useful in tests.
func NewNoop() *NoopLogger {
	return &NoopLogger{}
}

func (*NoopLogger) Trace(msg string, fields map[string]interface{}) {}
func (*NoopLogger) Debug(msg string, fields map[string]interface{}) {}
func (*NoopLogger) Info(msg string, fields map[string]interface{})  {}
func (*NoopLogger) Warn(msg string, fields map[string]interface{})  {}
func (*NoopLogger) Error(msg string, fields map[string]interface{}) {}
