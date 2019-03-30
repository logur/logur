package benchmarks

import (
	"io/ioutil"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/kitlogadapter"
)

func newKitlog() logur.Logger {
	logger := log.NewJSONLogger(ioutil.Discard)
	logger = level.NewFilter(logger, level.AllowAll())

	return kitlogadapter.New(logger)
}

func newDisabledKitlog() logur.Logger {
	logger := log.NewJSONLogger(ioutil.Discard)
	logger = level.NewFilter(logger, level.AllowError())

	return kitlogadapter.New(logger)
}
