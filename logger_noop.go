package logur

type noop struct{}

// NewNoop creates a no-op logger that can be used to silence
// all logging from this library. Also useful in tests.
func NewNoop() Logger {
	return &noop{}
}

func (*noop) Trace(msg string) {}
func (*noop) Debug(msg string) {}
func (*noop) Info(msg string)  {}
func (*noop) Warn(msg string)  {}
func (*noop) Error(msg string) {}

func (n *noop) WithFields(fields map[string]interface{}) Logger { return n }
