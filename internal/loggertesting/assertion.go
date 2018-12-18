package loggertesting

import (
	"testing"

	"github.com/goph/logur"
)

const (
	// Matches lines without matching the last newline character
	AllowNoNewLine uint8 = 1 << iota
)

// AssertLogEvents asserts that two LogEvents are identical.
func AssertLogEvents(t *testing.T, expected logur.LogEvent, actual logur.LogEvent) {
	t.Helper()

	if expected.Level != actual.Level {
		t.Errorf("expected log levels to be equal\ngot:  %s\nwant: %s", actual.Level, expected.Level)
	}

	if expected.Line != actual.Line {
		t.Errorf("expected log lines to be equal\ngot:  %q\nwant: %q", actual.Line, expected.Line)
	}

	if len(expected.Fields) != len(actual.Fields) {
		t.Errorf("expected log fields to be equal\ngot:  %v\nwant: %v", actual.Fields, expected.Fields)
	}

	for key, value := range expected.Fields {
		if actual.Fields[key] != value {
			t.Errorf("expected log fields to be equal\ngot:  %v\nwant: %v", actual.Fields, expected.Fields)

			break
		}
	}
}
