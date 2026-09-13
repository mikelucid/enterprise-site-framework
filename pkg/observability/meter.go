package observability

import "go.opentelemetry.io/otel"

func Meter(name string) any { return otel.GetMeterProvider().Meter(name) }
