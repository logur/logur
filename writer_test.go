package logur

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLevelWriter(t *testing.T) {
	logger := NewTestLogger()

	writer := NewLevelWriter(logger, ErrorLevel)
	defer writer.Close()

	_, err := fmt.Fprintln(writer, "message")
	if err != nil {
		t.Fatal("writing log event failed:", err.Error())
	}

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

	if event.Level != ErrorLevel {
		t.Errorf("expected level %q instead of %q", ErrorLevel.String(), event.Level.String())
	}

	if event.Line != "message" {
		t.Errorf("expected message \"message\" instead of %q", event.Line)
	}
}
