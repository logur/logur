package logur_test

import (
	"testing"

	. "logur.dev/logur"
	"logur.dev/logur/conformance"
	"logur.dev/logur/logtesting"
)

func TestFieldLogger(t *testing.T) {
	t.Run("WithFields", func(t *testing.T) {
		logger := &TestLogger{}
		l := WithFields(
			WithFields(
				WithFields(logger, map[string]interface{}{"key": "value"}),
				map[string]interface{}{"key": "value2"},
			),
			map[string]interface{}{"key": "value3"},
		)

		l.Info("message", map[string]interface{}{"key2": "value"})

		logEvent := LogEvent{
			Line:   "message",
			Level:  Info,
			Fields: map[string]interface{}{"key": "value3", "key2": "value"},
		}

		logtesting.AssertLogEventsEqual(t, logEvent, logEvent)
	})

	t.Run("WithField", func(t *testing.T) {
		logger := &TestLogger{}
		l := WithField(
			WithField(
				WithField(logger, "key", "value"),
				"key", "value2",
			),
			"key", "value3",
		)

		l.Info("message", map[string]interface{}{"key2": "value"})

		logEvent := LogEvent{
			Line:   "message",
			Level:  Info,
			Fields: map[string]interface{}{"key": "value3", "key2": "value"},
		}

		logtesting.AssertLogEventsEqual(t, logEvent, *logger.LastEvent())
	})

	suite := conformance.TestSuite{
		LoggerFactory: func(_ Level) (Logger, conformance.TestLogger) {
			logger := &TestLoggerFacade{}

			return WithFields(logger, map[string]interface{}{"key": "value"}), logger
		},
	}

	t.Run("Conformance", suite.Run)
}
