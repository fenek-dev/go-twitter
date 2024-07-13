package common

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// Init returns an instance of Jaeger Tracer.
func Init(ctx context.Context, service string) *sdktrace.TracerProvider {
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
	)
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatal("creating OTLP trace exporter: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(newResource(service)),
	)

	otel.SetTracerProvider(tp)

	// Set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp
}

func newResource(service string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(service),
		semconv.ServiceVersion("0.0.1"),
	)
}

func InjectMetadataIntoContext(ctx context.Context, metadata map[string]string) context.Context {
	propagator := otel.GetTextMapPropagator()

	return propagator.Extract(
		ctx,
		propagation.MapCarrier(metadata),
	)
}

func ExtractMetadataFromContext(ctx context.Context) map[string]string {
	propagator := otel.GetTextMapPropagator()

	metadata := map[string]string{}
	propagator.Inject(
		ctx,
		propagation.MapCarrier(metadata),
	)

	return metadata
}
