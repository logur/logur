package benchmarks

import (
	"io/ioutil"

	"github.com/goph/logur"
	"github.com/goph/logur/adapters/zapadapter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const _zap = "go.uber.org/zap"

func newZap() logur.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)

	return zapadapter.New(zap.New(zapcore.NewCore(enc, zapcore.AddSync(ioutil.Discard), zap.DebugLevel)).Sugar())
}


func newDisabledZap() logur.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)

	return zapadapter.New(zap.New(zapcore.NewCore(enc, zapcore.AddSync(ioutil.Discard), zap.ErrorLevel)).Sugar())
}
