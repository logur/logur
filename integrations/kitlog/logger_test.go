package kitlog

import (
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/testing"
)

func TestLogger_Log(t *testing.T) {
	testLogger := logur.NewTestLogger()
	logger := New(testLogger)

	_ = logger.Log("msg", "message", "key", "value")

	expected := logur.LogEvent{
		Line:  "message",
		Level: logur.Info,
		Fields: map[string]interface{}{
			"key": "value",
		},
	}

	logtesting.AssertLogEventsEqual(t, expected, *(testLogger.LastEvent()))
}

func TestLogger_Log_Level(t *testing.T) {
	tests := map[string]logur.Level{
		"trace":   logur.Trace,
		"debug":   logur.Debug,
		"info":    logur.Info,
		"warn":    logur.Warn,
		"warning": logur.Warn,
		"error":   logur.Error,
	}

	for level, llevel := range tests {
		level, llevel := level, llevel

		t.Run(level, func(t *testing.T) {
			testLogger := logur.NewTestLogger()
			logger := New(testLogger)
			_ = logger.Log("level", level)

			expected := logur.LogEvent{
				Level: llevel,
			}

			logtesting.AssertLogEventsEqual(t, expected, *(testLogger.LastEvent()))
		})
	}
}

func TestLogger_Log_MissingValue(t *testing.T) {
	testLogger := logur.NewTestLogger()
	logger := New(testLogger)

	_ = logger.Log("key")

	expected := logur.LogEvent{
		Level: logur.Info,
		Fields: map[string]interface{}{
			"key": "(MISSING)",
		},
	}

	logtesting.AssertLogEventsEqual(t, expected, *(testLogger.LastEvent()))
}
