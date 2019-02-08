package logrlog

import (
	"testing"

	"github.com/go-logr/logr"
)

func TestLogger(t *testing.T) {
	var logger interface{} = New(nil)

	if _, ok := logger.(logr.Logger); !ok {
		t.Error("Logger does not implement the logr.Logger interface")
	}
}
