package inmemory

import (
	"context"
	"github.com/google/uuid"
	pkg2 "github.com/opicaud/monorepo/events/eventstore/pkg"
	"github.com/opicaud/monorepo/events/pkg"
	v2beta "github.com/opicaud/monorepo/events/pkg/v2beta"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewInMemoryEventStore() *EventStore {
	build, _ := pkg2.NewEventStoreBuilder().Build("none")
	provider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return &EventStore{eventStore: build, tracer: otel.Tracer("non-grpc-inmemory-eventStore")}
}

func (f *EventStore) Save(ctx context.Context, events ...pkg.DomainEvent) (context.Context, []pkg.DomainEvent, error) {
	ctx, _ = f.startTrace(ctx, "Save")
	err := f.eventStore.Save(events...)
	return ctx, events, err
}

func (f *EventStore) Remove(ctx context.Context, uuid uuid.UUID) (context.Context, error) {
	return ctx, f.eventStore.Remove(uuid)
}

func (f *EventStore) Load(ctx context.Context, uuid uuid.UUID) (context.Context, []pkg.DomainEvent, error) {
	ctx, span := f.startTrace(ctx, "Load")
	load, err := f.eventStore.Load(uuid)
	if err != nil {
		span.RecordError(err)
	}
	return ctx, load, err
}
func (f *EventStore) startTrace(ctx context.Context, feature string) (context.Context, trace.Span) {
	ctx, span := f.tracer.Start(ctx, feature)
	defer span.End()
	return ctx, span
}

type EventStore struct {
	events     []pkg.DomainEvent
	eventStore v2beta.EventStore
	tracer     trace.Tracer
}

func (f *EventStore) GetHealthClient(ctx context.Context) (context.Context, grpc_health_v1.HealthClient) {
	return ctx, f.eventStore.GetHealthClient()
}
