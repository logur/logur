package logur

import (
	"testing"
	"time"
)

func TestNewStandardLogger(t *testing.T) {
	logger := NewTestLogger()
	stdLogger := NewStandardLogger(logger, Error, "", 0)

	const msg = "message"

	stdLogger.Println(msg)

	// Wait for the written data to reach the logger
	for i := 0; i < 3; i++ {
		if logger.Count() > 0 {
			break
		}

		time.Sleep(time.Duration((i+1)*10) * time.Millisecond)
	}

	if logger.Count() < 1 {
		t.Fatal("logger did not record any events")
	}

	event := logger.LastEvent()

	if event.Level != Error {
		t.Errorf("expected level %q instead of %q", Error.String(), event.Level.String())
	}

	if got, want := event.Line, msg; got != want {
		t.Errorf("expected message %q instead of %q", want, got)
	}
}
