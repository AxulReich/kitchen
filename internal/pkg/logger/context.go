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

func ToContext(ctx context.Context, l *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerContextKey, l)
}

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

func WithName(ctx context.Context, name string) context.Context {
	log := FromContext(ctx).Named(name)
	return ToContext(ctx, log)
}

func WithKV(ctx context.Context, key string, value interface{}) context.Context {
	log := FromContext(ctx).With(key, value)
	return ToContext(ctx, log)
}

func WithFields(ctx context.Context, fields ...zap.Field) context.Context {
	log := FromContext(ctx).
		Desugar().
		With(fields...).
		Sugar()
	return ToContext(ctx, log)
}
