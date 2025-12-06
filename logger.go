package zapl

import (
	"github.com/mandelsoft/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type noOutput struct {
}

var _ zapcore.WriteSyncer = (*noOutput)(nil)

func (s *noOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (s *noOutput) Sync() error {
	return nil
}

// NewStatic provides a zap logger for a given logging.Context.
// The returned logger logs into the logging.Context.
// The given name is used as sub realm for the base realm *zap*.
// If you want the other direction to use zap as base logger for a logging context
// use github.com/go-logr/zapr.NewLogger to obtain a logr.Logger for
// a zap logger.
func NewStatic(ctx logging.Context, name string, opts ...zap.Option) *zap.Logger {
	return New(ctx, name, StaticLogging, opts...)
}

// NewDynamic like NewStatic but used Unbound loggers.
// Unbound loggers consider the logging rules when used.
func NewDynamic(ctx logging.Context, name string, opts ...zap.Option) *zap.Logger {
	return New(ctx, name, DynamicLogging, opts...)
}

func New(ctx logging.Context, basename string, creator LoggerFactory, opts ...zap.Option) *zap.Logger {
	zl := ctx.Logger(logging.NewRealm("zap"))

	il := zapcore.DebugLevel
	if !zl.Enabled(logging.DebugLevel) {
		il = zapcore.InfoLevel
	}
	if !zl.Enabled(logging.InfoLevel) {
		il = zapcore.WarnLevel
	}
	if !zl.Enabled(logging.WarnLevel) {
		il = zapcore.ErrorLevel
	}
	lvl := zap.NewAtomicLevelAt(il)
	enc := NewEncoder(ctx, creator)

	zcore := zapcore.NewCore(enc, &noOutput{}, lvl)
	log := zap.New(zcore)
	if basename != "" {
		log = log.Named(basename)
	}
	return log
}
