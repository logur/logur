package invisionlog

import (
	"testing"

	log "github.com/InVisionApp/go-logger"
)

func TestLogger(t *testing.T) {
	var logger interface{} = New(nil)

	if _, ok := logger.(log.Logger); !ok {
		t.Error("Logger does not implement the log.Logger interface")
	}
}
