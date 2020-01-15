package logtesting

import (
	"testing"

	"logur.dev/logur"
)

// AssertLogEventsEqual asserts that two LogEvents are identical.
func AssertLogEventsEqual(t *testing.T, expected logur.LogEvent, actual logur.LogEvent) {
	t.Helper()

	err := actual.AssertEquals(expected)
	if err != nil {
		t.Errorf("%+v", err)
	}
}
