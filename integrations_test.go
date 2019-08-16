package logur_test

import (
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/logtesting"
)

func TestKitLogger_Log(t *testing.T) {
	testLogger := NewTestLogger()
	logger := NewKitLogger(testLogger)

	_ = logger.Log("msg", "message", "key", "value")

	expected := LogEvent{
		Line:  "message",
		Level: Info,
		Fields: map[string]interface{}{
			"key": "value",
		},
	}

	logtesting.AssertLogEventsEqual(t, expected, *(testLogger.LastEvent()))
}

func TestKitLogger_Log_Level(t *testing.T) {
	tests := map[string]Level{
		"trace":   Trace,
		"debug":   Debug,
		"info":    Info,
		"warn":    Warn,
		"warning": Warn,
		"error":   Error,
	}

	for level, llevel := range tests {
		level, llevel := level, llevel

		t.Run(level, func(t *testing.T) {
			testLogger := NewTestLogger()
			logger := NewKitLogger(testLogger)
			_ = logger.Log("level", level)

			expected := LogEvent{
				Level: llevel,
			}

			logtesting.AssertLogEventsEqual(t, expected, *(testLogger.LastEvent()))
		})
	}
}

func TestKitLogger_Log_MissingValue(t *testing.T) {
	testLogger := NewTestLogger()
	logger := NewKitLogger(testLogger)

	_ = logger.Log("key")

	expected := LogEvent{
		Level: Info,
		Fields: map[string]interface{}{
			"key": "(MISSING)",
		},
	}

	logtesting.AssertLogEventsEqual(t, expected, *(testLogger.LastEvent()))
}
