package logur_test

import (
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
)

func newTestLoggerSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
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

	loggertesting.AssertLogEvents(t, event, events[0])
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

	loggertesting.AssertLogEvents(t, event, *lastEvent)
}
