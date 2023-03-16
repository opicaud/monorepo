package pkg

import (
	"context"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/shape-app/events/eventstore/grpc/inmemory/internal"
	"github.com/opicaud/monorepo/shape-app/events/eventstore/grpc/proto/gen"
	"github.com/opicaud/monorepo/shape-app/events/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type InMemoryGrpcEventStore struct {
	builder GrpcBuilder
}

func NewInMemoryGrpcEventStore() *InMemoryGrpcEventStore {
	return new(InMemoryGrpcEventStore)
}

func (i *InMemoryGrpcEventStore) Save(events ...pkg.DomainEvent) error {
	err := i.builder.Connect().Save(events...)
	return err
}

func (i *InMemoryGrpcEventStore) Load(id uuid.UUID) ([]pkg.DomainEvent, error) {
	events, err := i.builder.Connect().Load(id)
	return events, err
}

type GrpcBuilder struct {
	conn   *grpc.ClientConn
	client gen.EventStoreClient
	ctx    context.Context
	cancel context.CancelFunc
	err    error
}

func (g *GrpcBuilder) Connect() *GrpcBuilder {
	g.conn, g.err = grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if g.err != nil {
		log.Panic(g.err)
	}

	g.ctx, g.cancel = context.WithTimeout(context.Background(), time.Second)
	g.client = gen.NewEventStoreClient(g.conn)
	return g
}

func (g *GrpcBuilder) deferred() {
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}(g.conn)
	defer g.cancel()
}

func (g *GrpcBuilder) Save(events ...pkg.DomainEvent) error {
	grpcEvents := g.grpcEvents(events)
	_, err := g.client.Save(g.ctx, grpcEvents)
	g.deferred()
	return err
}

func (g *GrpcBuilder) Load(uuid uuid.UUID) ([]pkg.DomainEvent, error) {
	id := gen.UUID{Id: uuid.String()}
	var events []pkg.DomainEvent
	response, err := g.client.Load(g.ctx, &id)
	if err == nil {
		events = domainEvents(response.Events.Event)
	}
	g.deferred()
	return events, err
}
func (g *GrpcBuilder) grpcEvents(events []pkg.DomainEvent) *gen.Events {
	grpcEvents := &gen.Events{}
	for _, event := range events {
		grpcEvents.Event = append(grpcEvents.Event, grpcEvent(event))
	}
	return grpcEvents
}

func grpcEvent(event pkg.DomainEvent) *gen.Event {
	return &gen.Event{
		AggregateId: &gen.UUID{Id: event.AggregateId().String()},
		Name:        event.Name(),
		Data:        event.Data(),
	}
}

func domainEvents(events []*gen.Event) []pkg.DomainEvent {
	var domainEvents []pkg.DomainEvent
	for _, event := range events {
		domainEvents = append(domainEvents, domainEvent(event))
	}
	return domainEvents
}

func domainEvent(event *gen.Event) pkg.DomainEvent {
	id, _ := uuid.Parse(event.AggregateId.Id)
	return eventstore.NewStandardEvent(id, event.GetName(), event.GetData())
}
