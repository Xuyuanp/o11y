package o11y

import (
	"context"

	"go.opentelemetry.io/contrib/exporters/autoexport"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTrace(ctx context.Context) (*sdktrace.TracerProvider, error) {
	exporter, err := autoexport.NewSpanExporter(ctx)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))

	otel.SetTracerProvider(tp)
	defaultPropgates := []propagation.TextMapPropagator{
		b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader)),
		propagation.Baggage{},
	}
	otel.SetTextMapPropagator(autoprop.NewTextMapPropagator(defaultPropgates...))

	return tp, nil
}
