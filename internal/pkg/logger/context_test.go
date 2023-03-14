package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestFromContext_GlobalLogger(t *testing.T) {
	logger := FromContext(context.Background())

	assert.Equal(t, global, logger)
}

func TestFromContext_WithLogger(t *testing.T) {
	l := New(zapcore.DebugLevel)
	ctx := ToContext(context.Background(), l)

	assert.Equal(t, l, FromContext(ctx))
}

func TestLoggerWithSpanContext(t *testing.T) {
	sc := jaeger.NewSpanContext(
		jaeger.TraceID{Low: 0x33328bc0eac485ce},
		jaeger.SpanID(0xa48b167265f65931),
		jaeger.SpanID(0x33328bc0eac485ce),
		true,
		nil,
	)

	buf := bytes.Buffer{}

	l := loggerWithSpanContext(loggerWithWriter(&buf), sc)
	l.Debug("hello world")

	var decoded map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "33328bc0eac485ce", decoded["trace_id"])
	assert.Equal(t, "a48b167265f65931", decoded["span_id"])
}

func TestLoggerWithName(t *testing.T) {
	buf := bytes.Buffer{}
	l := loggerWithWriter(&buf)

	ctx := ToContext(context.Background(), l)
	ctx = WithName(ctx, "test-logger")

	FromContext(ctx).Debug(ctx, "hello world")

	t.Logf("%s", &buf)
	var decoded map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test-logger", decoded["logger"])
}

func TestLoggerWithKV(t *testing.T) {
	buf := bytes.Buffer{}
	l := loggerWithWriter(&buf)

	ctx := ToContext(context.Background(), l)
	ctx = WithKV(ctx, "apples", 500)

	FromContext(ctx).Debug(ctx, "hello world")

	t.Logf("%s", &buf)
	var decoded map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 500, decoded["apples"])
}

func TestLoggerWithFields(t *testing.T) {
	buf := bytes.Buffer{}
	l := loggerWithWriter(&buf)

	ctx := ToContext(context.Background(), l)
	ctx = WithFields(ctx,
		zap.String("kafka-topic", "test-topic"),
		zap.Int32("kafka-partition", 420),
	)

	FromContext(ctx).Debug(ctx, "hello world")

	t.Logf("%s", &buf)
	var decoded map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, "test-topic", decoded["kafka-topic"])
	assert.EqualValues(t, 420, decoded["kafka-partition"])
}

func loggerWithWriter(w io.Writer) *zap.SugaredLogger {
	sink := zapcore.AddSync(w)
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
			sink,
			zap.NewAtomicLevelAt(zapcore.DebugLevel),
		),
	).Sugar()
}

func ExampleWithKV() {
	ctx := context.Background()
	ctx = WithKV(ctx, "my key", "my value")

	_ = ctx
}

func ExampleWithFields() {
	ctx := context.Background()
	ctx = WithFields(ctx,
		zap.String("kafka-topic", "my topic"),
		zap.Int32("kafka-partition", 1),
	)

	_ = ctx
}

func ExampleWithName() {
	ctx := context.Background()
	ctx = WithName(ctx, "GetApples")    // -> "GetApples"
	ctx = WithName(ctx, "AppleManager") // - > "GetApples.AppleManager"
	ctx = WithName(ctx, "DB")           // -> "GetApples.AppleManager.DB"

	_ = ctx
}
