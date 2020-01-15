package logur_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	. "logur.dev/logur"
	"logur.dev/logur/conformance"
	"logur.dev/logur/logtesting"
)

func TestLogEvents_Equals(t *testing.T) {
	tests := map[string]struct {
		expected LogEvent
		actual   LogEvent
	}{
		"level": {
			expected: LogEvent{
				Level: Trace,
			},
			actual: LogEvent{
				Level: Debug,
			},
		},
		"line": {
			expected: LogEvent{
				Line: "message",
			},
			actual: LogEvent{
				Line: "other message",
			},
		},
		"fields length": {
			expected: LogEvent{
				Fields: Fields{},
			},
			actual: LogEvent{
				Fields: Fields{"key1": "value1"},
			},
		},
		"fields value": {
			expected: LogEvent{
				Fields: Fields{"key1": "value1"},
			},
			actual: LogEvent{
				Fields: Fields{"key1": "value2"},
			},
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			if test.actual.Equals(test.expected) {
				t.Errorf(
					"log events should not be equal\nexpected: %+v\nactual:  %+v",
					test.expected,
					test.actual,
				)
			}
		})
	}
}

func TestLogEvents_AssertEquals(t *testing.T) {
	actual := LogEvent{
		Line:   "something happened",
		Level:  Trace,
		Fields: Fields{"key": "value"},
	}

	expected := LogEvent{
		Line:   "something else happened",
		Level:  Info,
		Fields: Fields{"key2": "value2"},
	}

	const expectedMessage = "failed to assert that log events are equal"
	const expectedVerboseMessage = `failed to assert that log events are equal
expected:
    line:   something else happened
    level:  info
    fields: map[key2:value2]
actual:
    line:   something happened
    level:  trace
    fields: map[key:value]
`

	err := actual.AssertEquals(expected)
	if err == nil {
		t.Fatal("assertion is expected to fail")
	}

	if want, have := expectedMessage, fmt.Sprintf("%s", err); want != have {
		t.Errorf("unexpexted error message: %v", have)
	}

	if want, have := expectedVerboseMessage, fmt.Sprintf("%+v", err); want != have {
		t.Errorf("unexpexted error message: %v", have)
	}
}

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

	suite := conformance.TestSuite{
		LoggerFactory: func(_ Level) (Logger, conformance.TestLogger) {
			logger := &TestLogger{}

			return logger, logger
		},
	}

	t.Run("Conformance", suite.Run)
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

	// T O D O: Conformance tests
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

	suite := conformance.TestSuite{
		LoggerFactory: func(_ Level) (Logger, conformance.TestLogger) {
			logger := &TestLoggerFacade{}

			return logger, logger
		},
	}

	t.Run("Conformance", suite.Run)
}
