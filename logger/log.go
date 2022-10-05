package logger

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	global       *zap.SugaredLogger
	defaultLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
)

func init() {
	SetLogger(New(defaultLevel))
	//metrics.MustRegister(messageCounters)
}

// InitFromConfig ...
func InitFromConfig(lvl string) {
	l := map[string]zapcore.Level{
		"DEBUG": zapcore.DebugLevel,
		"INFO":  zapcore.InfoLevel,
		"WARN":  zapcore.WarnLevel,
		"ERROR": zapcore.ErrorLevel,
	}
	if v, ok := l[lvl]; ok {
		SetLevel(v)
		return
	}
	SetLevel(defaultLevel.Level())
}

// New ...
func New(level zapcore.LevelEnabler, options ...zap.Option) *zap.SugaredLogger {
	return NewWithSink(level, os.Stdout, options...)
}

// NewWithSink ...
func NewWithSink(level zapcore.LevelEnabler, sink io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = defaultLevel
	}
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(sink),
			level,
		),
		options...,
	).Sugar()
}

// Logger ...
func Logger() *zap.SugaredLogger {
	return global
}

// SetLogger ...
func SetLogger(l *zap.SugaredLogger) {
	global = l
}

// Level ...
func Level() zapcore.Level {
	return defaultLevel.Level()
}

// SetLevel ...
func SetLevel(l zapcore.Level) {
	defaultLevel.SetLevel(l)
}

// Debug ...
func Debug(ctx context.Context, args ...interface{}) {
	//debugMessageCounter.Inc()
	FromContext(ctx).Debug(args...)
}

// Debugf ...
func Debugf(ctx context.Context, format string, args ...interface{}) {
	//debugMessageCounter.Inc()
	FromContext(ctx).Debugf(format, args...)
}

// DebugKV ...
func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	//debugMessageCounter.Inc()
	FromContext(ctx).Debugw(message, kvs...)
}

// Info ...
func Info(ctx context.Context, args ...interface{}) {
	//infoMessageCounter.Inc()
	FromContext(ctx).Info(args...)
}

// Infof ...
func Infof(ctx context.Context, format string, args ...interface{}) {
	//infoMessageCounter.Inc()
	FromContext(ctx).Infof(format, args...)
}

// InfoKV ...
func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	//infoMessageCounter.Inc()
	FromContext(ctx).Infow(message, kvs...)
}

// Warn ...
func Warn(ctx context.Context, args ...interface{}) {
	//warnMessageCounter.Inc()
	FromContext(ctx).Warn(args...)
}

// Warnf ...
func Warnf(ctx context.Context, format string, args ...interface{}) {
	//warnMessageCounter.Inc()
	FromContext(ctx).Warnf(format, args...)
}

// WarnKV ...
func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	//warnMessageCounter.Inc()
	FromContext(ctx).Warnw(message, kvs...)
}

// Error ...
func Error(ctx context.Context, args ...interface{}) {
	//errorMessageCounter.Inc()
	FromContext(ctx).Error(args...)
}

// Errorf ...
func Errorf(ctx context.Context, format string, args ...interface{}) {
	//errorMessageCounter.Inc()
	FromContext(ctx).Errorf(format, args...)
}

// ErrorKV ...
func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	//errorMessageCounter.Inc()
	FromContext(ctx).Errorw(message, kvs...)
}

// Fatal ...
func Fatal(ctx context.Context, args ...interface{}) {
	//fatalMessageCounter.Inc()
	FromContext(ctx).Fatal(args...)
}

// Fatalf ...
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	//fatalMessageCounter.Inc()
	FromContext(ctx).Fatalf(format, args...)
}

// FatalKV ...
func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	//fatalMessageCounter.Inc()
	FromContext(ctx).Fatalw(message, kvs...)
}

// Panic ...
func Panic(ctx context.Context, args ...interface{}) {
	//panicMessageCounter.Inc()
	FromContext(ctx).Panic(args...)
}

// Panicf ...
func Panicf(ctx context.Context, format string, args ...interface{}) {
	//panicMessageCounter.Inc()
	FromContext(ctx).Panicf(format, args...)
}

// PanicKV ...
func PanicKV(ctx context.Context, message string, kvs ...interface{}) {
	//panicMessageCounter.Inc()
	FromContext(ctx).Panicw(message, kvs...)
}
