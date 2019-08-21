package benchmarks

import (
	"io/ioutil"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	zapadapter "logur.dev/adapter/zap"

	"github.com/goph/logur"
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
