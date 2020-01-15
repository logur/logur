package logur_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	. "logur.dev/logur"
	"logur.dev/logur/conformance"
	"logur.dev/logur/logtesting"
)

func TestContextLogger(t *testing.T) {
	t.Run("NoFields", func(t *testing.T) {
		testLogger := &TestLoggerFacade{}

		logger := NewLoggerContext(testLogger, func(ctx context.Context) map[string]interface{} {
			return nil
		})

		logger.InfoContext(context.Background(), "message")

		logEvent := LogEvent{
			Line:  "message",
			Level: Info,
		}

		logtesting.AssertLogEventsEqual(t, logEvent, *testLogger.LastEvent())
	})

	t.Run("ContextFields", func(t *testing.T) {
		testLogger := &TestLoggerFacade{}

		logger := NewLoggerContext(testLogger, func(ctx context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key": "value",
			}
		})

		logger.InfoContext(context.Background(), "message")

		logEvent := LogEvent{
			Line:   "message",
			Level:  Info,
			Fields: map[string]interface{}{"key": "value"},
		}

		logtesting.AssertLogEventsEqual(t, logEvent, *testLogger.LastEvent())
	})

	t.Run("Fields", func(t *testing.T) {
		testLogger := &TestLoggerFacade{}

		logger := NewLoggerContext(testLogger, func(ctx context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}
		})

		logger.InfoContext(context.Background(), "message", map[string]interface{}{
			"key":  "another value",
			"key3": "value3",
		})

		logEvent := LogEvent{
			Line:  "message",
			Level: Info,
			Fields: map[string]interface{}{
				"key":  "another value",
				"key2": "value2",
				"key3": "value3",
			},
		}

		logtesting.AssertLogEventsEqual(t, logEvent, *testLogger.LastEvent())
	})

	t.Run("Conformance", func(t *testing.T) {
		t.Run("Logger", func(t *testing.T) {
			suite := conformance.TestSuite{
				LoggerFactory: func(_ Level) (Logger, conformance.TestLogger) {
					logger := &TestLogger{}

					return NewLoggerContext(logger, func(ctx context.Context) map[string]interface{} {
						return nil
					}), logger
				},
			}

			suite.Run(t)
		})

		t.Run("Facade", func(t *testing.T) {
			suite := conformance.TestSuite{
				LoggerFactory: func(_ Level) (Logger, conformance.TestLogger) {
					logger := &TestLoggerFacade{}

					return NewLoggerContext(logger, func(ctx context.Context) map[string]interface{} {
						return nil
					}), logger
				},
			}

			suite.Run(t)
		})
	})
}

func TestContextExtractors(t *testing.T) {
	extractor := ContextExtractors(
		func(_ context.Context) map[string]interface{} {
			return nil
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key":  "another_value",
				"key3": "value3",
			}
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key4": time.Minute,
			}
		},
	)

	expected := map[string]interface{}{
		"key":  "another_value",
		"key2": "value2",
		"key3": "value3",
		"key4": time.Minute,
	}

	if want, have := expected, extractor(context.Background()); !reflect.DeepEqual(want, have) {
		t.Errorf("unexpexted details\nexpected: %v\nactual:   %v", want, have)
	}
}
