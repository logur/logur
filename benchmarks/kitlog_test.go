package benchmarks

import (
	"io/ioutil"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitadapter "logur.dev/adapter/kit"

	"github.com/goph/logur"
)

func newKitlog() logur.Logger {
	logger := log.NewJSONLogger(ioutil.Discard)
	logger = level.NewFilter(logger, level.AllowAll())

	return kitadapter.New(logger)
}

func newDisabledKitlog() logur.Logger {
	logger := log.NewJSONLogger(ioutil.Discard)
	logger = level.NewFilter(logger, level.AllowError())

	return kitadapter.New(logger)
}
