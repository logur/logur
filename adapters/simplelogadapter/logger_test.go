package simplelogadapter

import (
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
)

type testLogger struct {
	*logur.TestLogger
}

func newTestLogger() *testLogger {
	return &testLogger{logur.NewTestLogger()}
}

func (l *testLogger) WithFields(fields Fields) Logger {
	return &testLogger{l.TestLogger.WithFields(logur.Fields(fields)).(*logur.TestLogger)}
}

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		LogEventAssertionFlags: 0 | loggertesting.SkipRawLine | loggertesting.AllowNoNewLine,
		LoggerFactory: func() (logur.Logger, func() []logur.LogEvent) {
			logger := newTestLogger()

			return New(logger), func() []logur.LogEvent { // nolint: gocritic
				return logger.Events()
			}
		},
	}
}

func TestLogger_Levels(t *testing.T) {
	newTestSuite().TestLevels(t)
}

func TestLogger_Levelsln(t *testing.T) {
	newTestSuite().TestLevelsln(t)
}

func TestLogger_Levelsf(t *testing.T) {
	newTestSuite().TestLevelsf(t)
}
