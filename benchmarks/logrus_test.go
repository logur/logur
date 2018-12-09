package benchmarks

import (
	"io/ioutil"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/logrusadapter"
	"github.com/sirupsen/logrus"
)

const logrusPackage = "github.com/sirupsen/logrus"

func newDisabledLogrus() logur.Logger {
	logger := logrus.New()
	logger.Level = logrus.ErrorLevel
	logger.Out = ioutil.Discard

	return logrusadapter.New(logger)
}

func newLogrus() logur.Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = ioutil.Discard

	return logrusadapter.New(logger)
}
