package benchmarks

import (
	"io/ioutil"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/logrusadapter"
	"github.com/sirupsen/logrus"
)

const _logrus = "github.com/sirupsen/logrus"

func newLogrus() logur.Logger {
	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logger.Out = ioutil.Discard

	return logrusadapter.New(logger)
}


func newDisabledLogrus() logur.Logger {
	logger := logrus.New()
	logger.Level = logrus.ErrorLevel
	logger.Out = ioutil.Discard

	return logrusadapter.New(logger)
}
