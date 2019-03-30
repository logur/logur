package benchmarks

import (
	"io/ioutil"

	"github.com/rs/zerolog"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/zerologadapter"
)

func newZerolog() logur.Logger {
	logger := zerolog.New(ioutil.Discard)
	logger.Level(zerolog.DebugLevel)

	return zerologadapter.New(logger)
}

func newDisabledZerolog() logur.Logger {
	logger := zerolog.New(ioutil.Discard)
	logger.Level(zerolog.ErrorLevel)

	return zerologadapter.New(logger)
}
