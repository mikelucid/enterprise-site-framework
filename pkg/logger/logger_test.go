package logger

import "testing"

func TestContextHelpers(t *testing.T) {
	ctx := WithRequestID(WithTraceID(t.Context(), "trace-1"), "req-1")
	if got, ok := TraceIDFromContext(ctx); !ok || got != "trace-1" {
		t.Fatalf("unexpected trace: %q", got)
	}
	if got, ok := RequestIDFromContext(ctx); !ok || got != "req-1" {
		t.Fatalf("unexpected request id: %q", got)
	}
}
