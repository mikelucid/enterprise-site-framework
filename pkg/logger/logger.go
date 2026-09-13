package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level string `mapstructure:"level"`
}

type Logger struct{ *zap.Logger }

func New(cfg Config) (*Logger, error) {
	level := zapcore.InfoLevel
	_ = level.UnmarshalText([]byte(cfg.Level))
	zcfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zl, err := zcfg.Build()
	if err != nil {
		return nil, err
	}
	return &Logger{Logger: zl}, nil
}

func (l *Logger) WithTrace(ctx context.Context) *zap.Logger {
	fields := []zap.Field{}
	if tid, ok := TraceIDFromContext(ctx); ok {
		fields = append(fields, zap.String("trace_id", tid))
	}
	if rid, ok := RequestIDFromContext(ctx); ok {
		fields = append(fields, zap.String("request_id", rid))
	}
	return l.With(fields...)
}

type contextKey string

const (
	traceIDKey   contextKey = "trace_id"
	requestIDKey contextKey = "request_id"
)

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func TraceIDFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(traceIDKey).(string)
	return v, ok
}

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func RequestIDFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(requestIDKey).(string)
	return v, ok
}
