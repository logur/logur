package logtesting

import (
	"testing"

	"github.com/goph/logur"
)

// AssertLogEvents asserts that two LogEvents are identical.
func AssertLogEvents(t *testing.T, expected logur.LogEvent, actual logur.LogEvent) {
	t.Helper()

	err := logur.AssertLogEventsEqual(expected, actual)
	if err != nil {
		t.Error(err)
	}
}
