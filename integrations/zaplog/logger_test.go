package zaplog

import (
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/logtesting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestWith(t *testing.T) {
	logger := logur.NewTestLogger()
	log := New(logger)
	log.With(zap.String("foo", "bar")).Info("hello", zap.String("baz", "quux"))

	logEvent := logur.LogEvent{
		Line:  "hello",
		Level: logur.Info,
		Fields: map[string]interface{}{
			"foo": "bar",
			"baz": "quux",
		},
	}

	logtesting.AssertLogEventsEqual(t, logEvent, *(logger.LastEvent()))
}

func TestLogger(t *testing.T) {
	const msg = "hello"

	logger := logur.NewTestLogger()
	log := New(logger)

	tests := []struct {
		f     func(string, ...zapcore.Field)
		level zapcore.Level
		want  logur.Level
	}{
		{log.Debug, zapcore.DebugLevel, logur.Debug},
		{log.Info, zapcore.InfoLevel, logur.Info},
		{log.Warn, zapcore.WarnLevel, logur.Warn},
		{log.Error, zapcore.ErrorLevel, logur.Error},
		{log.DPanic, zapcore.DPanicLevel, logur.Error},
	}

	for _, test := range tests {
		test := test

		t.Run(test.want.String(), func(t *testing.T) {
			test.f(msg)

			logEvent := logur.LogEvent{
				Line:  msg,
				Level: test.want,
			}

			logtesting.AssertLogEventsEqual(t, logEvent, *(logger.LastEvent()))

			if ce := log.Check(test.level, msg); ce != nil {
				ce.Write()
			}

			logtesting.AssertLogEventsEqual(t, logEvent, *(logger.LastEvent()))
		})
	}
}
