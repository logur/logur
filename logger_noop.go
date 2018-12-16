package logur

type noop struct{}

// NewNoop creates a no-op logger that can be used to silence
// all logging from this library. Also useful in tests.
func NewNoop() Logger {
	return &noop{}
}

func (*noop) Trace(args ...interface{}) {}
func (*noop) Debug(args ...interface{}) {}
func (*noop) Info(args ...interface{})  {}
func (*noop) Warn(args ...interface{})  {}
func (*noop) Error(args ...interface{}) {}

func (*noop) Traceln(args ...interface{}) {}
func (*noop) Debugln(args ...interface{}) {}
func (*noop) Infoln(args ...interface{})  {}
func (*noop) Warnln(args ...interface{})  {}
func (*noop) Errorln(args ...interface{}) {}

func (*noop) Tracef(format string, args ...interface{}) {}
func (*noop) Debugf(format string, args ...interface{}) {}
func (*noop) Infof(format string, args ...interface{})  {}
func (*noop) Warnf(format string, args ...interface{})  {}
func (*noop) Errorf(format string, args ...interface{}) {}

func (n *noop) WithFields(fields map[string]interface{}) Logger { return n }
