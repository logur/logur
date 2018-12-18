package logur

type noop struct{}

// NewNoop creates a no-op logger that can be used to silence
// all logging from this library. Also useful in tests.
func NewNoop() Logger {
	return &noop{}
}

func (*noop) Trace(msg string, fields map[string]interface{}) {}
func (*noop) Debug(msg string, fields map[string]interface{}) {}
func (*noop) Info(msg string, fields map[string]interface{})  {}
func (*noop) Warn(msg string, fields map[string]interface{})  {}
func (*noop) Error(msg string, fields map[string]interface{}) {}
