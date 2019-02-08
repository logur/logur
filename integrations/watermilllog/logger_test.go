package watermilllog

import (
	"testing"

	"github.com/ThreeDotsLabs/watermill"
)

func TestLogger(t *testing.T) {
	var logger interface{} = New(nil)

	if _, ok := logger.(watermill.LoggerAdapter); !ok {
		t.Error("Logger does not implement the watermill.LoggerAdapter interface")
	}
}
