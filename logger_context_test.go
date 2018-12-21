package logur_test

import (
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/testing"
)

func newContextualLoggerTestSuite() *logtesting.LoggerTestSuite {
	return &logtesting.LoggerTestSuite{
		LoggerFactory: func(level Level) (Logger, func() []LogEvent) {
			logger := NewTestLogger()

			return WithFields(logger, map[string]interface{}{"key": "value"}), logger.Events
		},
	}
}

func TestContextualLogger_Levels(t *testing.T) {
	newContextualLoggerTestSuite().TestLevels(t)
}

func TestWithFields(t *testing.T) {
	logger := NewTestLogger()
	ctxlogger := WithFields(
		WithFields(
			WithFields(logger, map[string]interface{}{"key": "value"}),
			map[string]interface{}{"key": "value2"},
		),
		map[string]interface{}{"key": "value3"},
	)

	ctxlogger.Info("message", map[string]interface{}{"key2": "value"})

	logEvent := LogEvent{
		Line:   "message",
		Level:  Info,
		Fields: map[string]interface{}{"key": "value3", "key2": "value"},
	}

	logtesting.AssertLogEvents(t, logEvent, logEvent)
}
