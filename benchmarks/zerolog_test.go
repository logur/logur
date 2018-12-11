package benchmarks

import (
	"io/ioutil"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/zerologadapter"
	"github.com/rs/zerolog"
)

const _zerolog = "github.com/rs/zerolog"

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
