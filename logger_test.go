package logur_test

import (
	"context"
	"testing"

	. "logur.dev/logur"
	"logur.dev/logur/logtesting"
)

func TestLoggerToKV(t *testing.T) {
	t.Parallel()

	t.Run("KV", func(t *testing.T) {
		t.Parallel()

		for _, level := range Levels() {
			level := level

			t.Run(level.String(), func(t *testing.T) {
				t.Parallel()

				testLogger := &TestLoggerFacade{}

				logger := LoggerToKV(testLogger)

				KVLevelFunc(logger, level)("msg", "key", "value")

				expected := LogEvent{
					Line:  "msg",
					Level: level,
					Fields: map[string]interface{}{
						"key": "value",
					},
				}

				logtesting.AssertLogEventsEqual(t, expected, *(testLogger.LastEvent()))
			})
		}
	})

	t.Run("Context", func(t *testing.T) {
		t.Parallel()

		for _, level := range Levels() {
			level := level

			t.Run(level.String(), func(t *testing.T) {
				t.Parallel()

				testLogger := &TestLoggerFacade{}

				logger := LoggerToKV(testLogger)

				KVLevelContextFunc(logger, level)(context.Background(), "msg", "key", "value")

				expected := LogEvent{
					Line:  "msg",
					Level: level,
					Fields: map[string]interface{}{
						"key": "value",
					},
				}

				logtesting.AssertLogEventsEqual(t, expected, *(testLogger.LastEvent()))
			})
		}
	})
}
