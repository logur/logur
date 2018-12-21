package logur_test

import (
	"strings"
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/testing"
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

	err := AssertLogEventsEqual(event1, event2)
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
				Fields: map[string]interface{}{"key1": "value1"},
			},
			actual: LogEvent{
				Fields: map[string]interface{}{"key1": "value1", "key2": "value2"},
			},
			expectedError: "expected log fields to be equal\ngot:  map[key1:value1 key2:value2]\nwant: map[key1:value1]",
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
			err := AssertLogEventsEqual(test.expected, test.actual)

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
		LoggerFactory: func(level Level) (Logger, func() []LogEvent) {
			logger := NewTestLogger()
			return logger, func() []LogEvent { // nolint: gocritic
				return logger.Events()
			}
		},
	}
}

func TestTestLogger_Levels(t *testing.T) {
	newTestLoggerSuite().TestLevels(t)
}

func TestTestLogger_Count(t *testing.T) {
	logger := NewTestLogger()

	logger.Debug("message", nil)

	if got, want := logger.Count(), 1; got != want {
		t.Errorf("expected log event count to be %d, got %d", want, got)
	}
}

func TestTestLogger_Events(t *testing.T) {
	logger := NewTestLogger()

	logger.Debug("message", nil)

	events := logger.Events()

	if got, want := len(events), 1; got != want {
		t.Fatalf("expected log event count to be %d, got %d", want, got)
	}

	event := LogEvent{
		Level: Debug,
		Line:  "message",
	}

	logtesting.AssertLogEvents(t, event, events[0])
}

func TestTestLogger_LastEvent(t *testing.T) {
	logger := NewTestLogger()

	logger.Debug("message", nil)
	logger.Info("another message", nil)

	lastEvent := logger.LastEvent()

	if lastEvent == nil {
		t.Fatal("failed to get last event")
	}

	event := LogEvent{
		Level: Info,
		Line:  "another message",
	}

	logtesting.AssertLogEvents(t, event, *lastEvent)
}
