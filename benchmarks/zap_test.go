package benchmarks

import (
	"io/ioutil"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/zapadapter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZap() logur.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)

	return zapadapter.New(zap.New(zapcore.NewCore(enc, zapcore.AddSync(ioutil.Discard), zap.DebugLevel)))
}

func newDisabledZap() logur.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)

	return zapadapter.New(zap.New(zapcore.NewCore(enc, zapcore.AddSync(ioutil.Discard), zap.ErrorLevel)))
}
