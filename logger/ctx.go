package logger

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
)

type contextKey int

const (
	loggerContextKey contextKey = iota
)

// ToContext ...
func ToContext(ctx context.Context, l *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerContextKey, l)
}

// FromContext ...
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l := global

	if logger, ok := ctx.Value(loggerContextKey).(*zap.SugaredLogger); ok {
		l = logger
	}

	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return l
	}

	return loggerWithSpanContext(l, span.Context())
}

func loggerWithSpanContext(l *zap.SugaredLogger, sc opentracing.SpanContext) *zap.SugaredLogger {
	if sc, ok := sc.(jaeger.SpanContext); ok {
		return l.Desugar().With(
			zap.Stringer("trace_id", sc.TraceID()),
			zap.Stringer("span_id", sc.SpanID()),
		).Sugar()
	}

	return l
}
