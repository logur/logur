package logur_test

import (
	"testing"
	"time"

	. "logur.dev/logur"
	"logur.dev/logur/logtesting"
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

	event := LogEvent{
		Level: Error,
		Line:  msg,
	}

	logtesting.AssertLogEventsEqual(t, event, *(logger.LastEvent()))
}

func TestNewStandardErrorLogger(t *testing.T) {
	logger := NewTestLogger()
	stdLogger := NewErrorStandardLogger(logger, "", 0)

	const msg = "message"

	stdLogger.Println(msg)

	// Wait for the written data to reach the logger
	for i := 0; i < 3; i++ {
		if logger.Count() > 0 {
			break
		}

		time.Sleep(time.Duration((i+1)*10) * time.Millisecond)
	}

	event := LogEvent{
		Level: Error,
		Line:  msg,
	}

	logtesting.AssertLogEventsEqual(t, event, *(logger.LastEvent()))
}
