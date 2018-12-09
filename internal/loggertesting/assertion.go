package loggertesting

import (
	"reflect"
	"testing"

	"github.com/goph/logur"
)

// AssertLogEvents asserts that two LogEvents are identical.
func AssertLogEvents(t *testing.T, expected logur.LogEvent, actual logur.LogEvent, skipRawLine bool) {
	t.Helper()

	if expected.Level != actual.Level {
		t.Errorf("expected log levels to be equal\ngot:  %s\nwant: %s", expected.Level, actual.Level)
	}

	if expected.Line != actual.Line {
		t.Errorf("expected log lines to be equal\ngot:  %s\nwant: %s", expected.Line, actual.Line)
	}

	if !skipRawLine {
		if !reflect.DeepEqual(expected.RawLine, actual.RawLine) {
			t.Errorf("expected raw log lines to be equal\ngot:  %v\nwant: %v", expected.RawLine, actual.RawLine)
		}
	}

	if !reflect.DeepEqual(expected.Fields, actual.Fields) {
		t.Errorf("expected log fields to be equal\ngot:  %v\nwant: %v", expected.Fields, actual.Fields)
	}
}
