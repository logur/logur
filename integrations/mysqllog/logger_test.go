package mysqllog

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/goph/logur"
	"github.com/goph/logur/testing"
)

func TestLogger(t *testing.T) {
	var _ mysql.Logger = New(logur.NewNoopLogger())
}

func TestLogger_Print(t *testing.T) {
	testLogger := logur.NewTestLogger()
	logger := New(testLogger)

	logger.Print("message", 1, "message", 2)

	event := logur.LogEvent{
		Level: logur.Error,
		Line:  "message1message2",
	}

	logtesting.AssertLogEvents(t, event, *(testLogger.LastEvent()))
}
