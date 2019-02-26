package zaplog

import (
	"fmt"

	"github.com/goph/logur"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New returns a new zap logger.
func New(logger logur.Logger) *zap.Logger {
	z := &zapper{
		logger: logger,
	}

	if levelEnabler, ok := logger.(logur.LevelEnabler); ok {
		z.levelEnabler = levelEnabler
	}

	return zap.New(z)
}

// zapper is created based on github.com/uber-common/bark/zbark.zapper.
type zapper struct {
	logger       logur.Logger
	levelEnabler logur.LevelEnabler
}

func (z *zapper) Enabled(lvl zapcore.Level) bool {
	if z.levelEnabler == nil {
		return true
	}

	// Logur does not have fatal and panic levels,
	// but since they are above error in severity, we fall back to error
	if lvl > zapcore.ErrorLevel {
		lvl = zapcore.ErrorLevel
	}

	return z.levelEnabler.LevelEnabled(logur.Level(lvl + 2))
}

func (z *zapper) With(fs []zapcore.Field) zapcore.Core {
	me := zapcore.NewMapObjectEncoder()
	for _, f := range fs {
		f.AddTo(me)
	}

	return &zapper{logur.WithFields(z.logger, me.Fields), z.levelEnabler}
}

func (z *zapper) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(ent.Level) {
		return ce.AddCore(ent, z)
	}

	return ce
}

func (z *zapper) Write(ent zapcore.Entry, fs []zapcore.Field) error {
	me := zapcore.NewMapObjectEncoder()
	for _, f := range fs {
		f.AddTo(me)
	}

	var logFunc func(string, ...map[string]interface{})
	switch ent.Level {
	case zapcore.DebugLevel:
		logFunc = z.logger.Debug
	case zapcore.InfoLevel:
		logFunc = z.logger.Info
	case zapcore.WarnLevel:
		logFunc = z.logger.Warn
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		logFunc = z.logger.Error
	default:
		return fmt.Errorf("logur-to-zap compatibility wrapper got unknown level %v", ent.Level)
	}

	logFunc(ent.Message, me.Fields)

	return nil
}

func (z *zapper) Sync() error {
	// Logur doesn't expose a way to flush buffered messages.
	return nil
}
