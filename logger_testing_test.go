package logur_test

import (
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
)

func newTestLoggerSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		LoggerFactory: func() (Logger, func() []LogEvent) {
			logger := NewTestLogger()
			return logger, func() []LogEvent { // nolint: unlambda
				return logger.Events()
			}
		},
	}
}

func TestTestLogger_Levels(t *testing.T) {
	newTestLoggerSuite().TestLevels(t)
}

func TestTestLogger_Levelsln(t *testing.T) {
	newTestLoggerSuite().TestLevelsln(t)
}

func TestTestLogger_Levelsf(t *testing.T) {
	newTestLoggerSuite().TestLevelsf(t)
}

func TestTestLogger_Count(t *testing.T) {
	logger := NewTestLogger()

	logger.Debug("message")

	if got, want := logger.Count(), 1; got != want {
		t.Errorf("expected log event count to be %d, got %d", want, got)
	}
}

func TestTestLogger_Events(t *testing.T) {
	logger := NewTestLogger()

	logger.Debug("message")

	events := logger.Events()

	if got, want := len(events), 1; got != want {
		t.Fatalf("expected log event count to be %d, got %d", want, got)
	}

	event := LogEvent{
		Level:   DebugLevel,
		Line:    "message",
		RawLine: []interface{}{"message"},
	}

	loggertesting.AssertLogEvents(t, event, events[0], true)
}

func TestTestLogger_LastEvent(t *testing.T) {
	logger := NewTestLogger()

	logger.Debug("message")
	logger.Info("another message")

	lastEvent := logger.LastEvent()

	if lastEvent == nil {
		t.Fatal("failed to get last event")
	}

	event := LogEvent{
		Level:   InfoLevel,
		Line:    "another message",
		RawLine: []interface{}{"another message"},
	}

	loggertesting.AssertLogEvents(t, event, *lastEvent, true)
}
