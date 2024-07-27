package common_test

import (
	"context"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
	"testing"
)

func NewTestTracer(t *testing.T) (trace.Tracer, func(ctx context.Context) error) {
	t.Helper()
	provider := sdktrace.NewTracerProvider(sdktrace.WithBatcher(tracetest.NewNoopExporter()))

	return provider.Tracer("test"), provider.Shutdown
}
