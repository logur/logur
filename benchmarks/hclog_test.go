package benchmarks

import (
	"io/ioutil"

	"github.com/hashicorp/go-hclog"
	hclogadapter "logur.dev/adapter/hclog"

	"logur.dev/logur"
)

func newHclog() logur.Logger {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:  hclog.Trace,
		Output: ioutil.Discard,
	})

	return hclogadapter.New(logger)
}

func newDisabledHclog() logur.Logger {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:  hclog.Error,
		Output: ioutil.Discard,
	})

	return hclogadapter.New(logger)
}
