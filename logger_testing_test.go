package logur_test

import (
	"context"
	"strings"
	"testing"

	. "logur.dev/logur"
	"logur.dev/logur/logtesting"
)

func TestAssertLogEventsEqual(t *testing.T) {
	event1 := LogEvent{
		Level:  Trace,
		Line:   "message",
		Fields: map[string]interface{}{"key1": "value1", "key2": "value2"},
	}

	event2 := LogEvent{
		Level:  Trace,
		Line:   "message",
		Fields: map[string]interface{}{"key2": "value2", "key1": "value1"},
	}

	err := LogEventsEqual(event1, event2)
	if err != nil {
		t.Errorf("failed to assert that two identical event are equal: %s", strings.Replace(err.Error(), "\n", `\n`, -1))
	}
}

func TestAssertLogEventsEqual_Errors(t *testing.T) {
	tests := map[string]struct {
		expected      LogEvent
		actual        LogEvent
		expectedError string
	}{
		"level": {
			expected: LogEvent{
				Level: Trace,
			},
			actual: LogEvent{
				Level: Debug,
			},
			expectedError: "expected log levels to be equal\ngot:  debug\nwant: trace",
		},
		"line": {
			expected: LogEvent{
				Line: "message",
			},
			actual: LogEvent{
				Line: "other message",
			},
			expectedError: "expected log lines to be equal\ngot:  \"other message\"\nwant: \"message\"",
		},
		"fields length": {
			expected: LogEvent{
				Fields: map[string]interface{}{},
			},
			actual: LogEvent{
				Fields: map[string]interface{}{"key1": "value1"},
			},
			expectedError: "expected log fields to be equal\ngot:  map[key1:value1]\nwant: map[]",
		},
		"fields value": {
			expected: LogEvent{
				Fields: map[string]interface{}{"key1": "value1"},
			},
			actual: LogEvent{
				Fields: map[string]interface{}{"key1": "value2"},
			},
			expectedError: "expected log fields to be equal\ngot:  map[key1:value2]\nwant: map[key1:value1]",
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			err := LogEventsEqual(test.expected, test.actual)

			if err.Error() != test.expectedError {
				actualError := strings.Replace(err.Error(), "\n", `\n`, -1)
				expectedError := strings.Replace(test.expectedError, "\n", `\n`, -1)

				t.Errorf("expected log levels to be equal\ngot:  %s\nwant: %s", actualError, expectedError)
			}
		})
	}
}

func newTestLoggerSuite() *logtesting.LoggerTestSuite {
	return &logtesting.LoggerTestSuite{
		LoggerFactory: func(_ Level) (Logger, func() []LogEvent) {
			logger := &TestLogger{}
			return logger, func() []LogEvent { // nolint: gocritic
				return logger.Events()
			}
		},
	}
}

func TestTestLogger_Levels(t *testing.T) {
	newTestLoggerSuite().TestLevels(t)
}

func TestTestLogger(t *testing.T) {
	logger := &TestLogger{}

	logger.Debug("debug message")
	logger.Info("info message", map[string]interface{}{"key": "value"})

	if want, have := 2, logger.Count(); want != have {
		t.Errorf("unexpexted log event count\nexpected: %v\nactual:   %v", want, have)
	}

	events := logger.Events()

	if want, have := 2, len(events); want != have {
		t.Errorf("unexpexted returned log event count\nexpected: %v\nactual:   %v", want, have)
	}

	lastEvent := LogEvent{
		Level:  Info,
		Line:   "info message",
		Fields: map[string]interface{}{"key": "value"},
	}

	logtesting.AssertLogEventsEqual(t, LogEvent{Level: Debug, Line: "debug message"}, events[0])
	logtesting.AssertLogEventsEqual(t, lastEvent, events[1])
	logtesting.AssertLogEventsEqual(t, lastEvent, *logger.LastEvent())
}

func TestTestLoggerContext(t *testing.T) {
	logger := &TestLoggerContext{}

	logger.DebugContext(context.Background(), "debug message")
	logger.InfoContext(context.Background(), "info message", map[string]interface{}{"key": "value"})

	if want, have := 2, logger.Count(); want != have {
		t.Errorf("unexpexted log event count\nexpected: %v\nactual:   %v", want, have)
	}

	events := logger.Events()

	if want, have := 2, len(events); want != have {
		t.Errorf("unexpexted returned log event count\nexpected: %v\nactual:   %v", want, have)
	}

	lastEvent := LogEvent{
		Level:  Info,
		Line:   "info message",
		Fields: map[string]interface{}{"key": "value"},
	}

	logtesting.AssertLogEventsEqual(t, LogEvent{Level: Debug, Line: "debug message"}, events[0])
	logtesting.AssertLogEventsEqual(t, lastEvent, events[1])
	logtesting.AssertLogEventsEqual(t, lastEvent, *logger.LastEvent())
}

func TestTestLoggerFacade(t *testing.T) {
	logger := &TestLoggerFacade{}

	logger.Debug("debug message")
	logger.DebugContext(context.Background(), "another debug message")
	logger.InfoContext(context.Background(), "another info message", map[string]interface{}{"key": "value"})

	if want, have := 3, logger.Count(); want != have {
		t.Errorf("unexpexted log event count\nexpected: %v\nactual:   %v", want, have)
	}

	events := logger.Events()

	if want, have := 3, len(events); want != have {
		t.Errorf("unexpexted returned log event count\nexpected: %v\nactual:   %v", want, have)
	}

	lastEvent := LogEvent{
		Level:  Info,
		Line:   "another info message",
		Fields: map[string]interface{}{"key": "value"},
	}

	logtesting.AssertLogEventsEqual(t, LogEvent{Level: Debug, Line: "debug message"}, events[0])
	logtesting.AssertLogEventsEqual(t, LogEvent{Level: Debug, Line: "another debug message"}, events[1])
	logtesting.AssertLogEventsEqual(t, lastEvent, events[2])
	logtesting.AssertLogEventsEqual(t, lastEvent, *logger.LastEvent())
}
