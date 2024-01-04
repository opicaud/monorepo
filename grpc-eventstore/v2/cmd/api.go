package pkg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/cqrs/v3/pkg"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewInMemoryEventStoreWithoutGrpc() *InMemoryEventStoreWithoutGrpc {
	provider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return &InMemoryEventStoreWithoutGrpc{tracer: otel.Tracer("non-grpc-inmemory-eventStore")}
}

func (f *InMemoryEventStoreWithoutGrpc) Save(ctx context.Context, events ...pkg.DomainEvent) (context.Context, []pkg.DomainEvent, error) {
	ctx, _ = f.startTrace(ctx, "Save")
	f.events = append(f.events, events...)
	return ctx, events, nil
}

func (f *InMemoryEventStoreWithoutGrpc) Remove(ctx context.Context, uuid uuid.UUID) (context.Context, error) {
	f.events = []pkg.DomainEvent{}
	return ctx, nil

}

func (f *InMemoryEventStoreWithoutGrpc) Load(ctx context.Context, uuid uuid.UUID) (context.Context, []pkg.DomainEvent, error) {
	ctx, span := f.startTrace(ctx, "Load")
	w := 0
	for _, e := range f.events {
		if e.AggregateId() == uuid {
			f.events[w] = e
			w++
		}
	}
	if len(f.events[0:w]) == 0 {
		err := fmt.Errorf("No aggregate with Id %s has been found", uuid)
		return ctx, nil, err
		span.RecordError(err)
	}
	return ctx, f.events[0:w], nil
}
func (f *InMemoryEventStoreWithoutGrpc) startTrace(ctx context.Context, feature string) (context.Context, trace.Span) {
	ctx, span := f.tracer.Start(ctx, feature)
	defer span.End()
	return ctx, span
}

type InMemoryEventStoreWithoutGrpc struct {
	events []pkg.DomainEvent
	tracer trace.Tracer
}

func (f *InMemoryEventStoreWithoutGrpc) GetHealthClient(ctx context.Context) (context.Context, grpc_health_v1.HealthClient) {
	fmt.Println("you are using an inmemory eventstore, serving by default")
	return ctx, nil
}
