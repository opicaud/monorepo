package grpc

import (
	"context"
	"example2/domain/adapter"
	ac "example2/infra/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type StandardEvent struct {
	id uuid.UUID
}

func newStandardEvent() StandardEvent {
	event := StandardEvent{}
	event.id = uuid.New()
	return event
}

func (e StandardEvent) AggregateId() uuid.UUID {
	return e.id
}

type InMemoryGrpcEventStore struct {
	builder GrpcBuilder
}

func NewInMemoryGrpcEventStore() *InMemoryGrpcEventStore {
	return new(InMemoryGrpcEventStore)
}

func (i *InMemoryGrpcEventStore) Save(events ...adapter.DomainEvent) error {
	err := i.builder.Connect().Save(events...)
	return err
}

func (i *InMemoryGrpcEventStore) Load(id uuid.UUID) ([]adapter.DomainEvent, error) {
	events, err := i.builder.Connect().Load(id)
	return events, err
}

type GrpcBuilder struct {
	conn   grpc.ClientConn
	client ac.EventStoreClient
	ctx    context.Context
	cancel context.CancelFunc
}

func (g *GrpcBuilder) Connect() *GrpcBuilder {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	g.conn = *conn
	g.ctx, g.cancel = context.WithTimeout(context.Background(), time.Second)
	g.client = ac.NewEventStoreClient(&g.conn)
	return g
}

func (g *GrpcBuilder) deferred() {
	defer g.conn.Close()
	defer g.cancel()
}

func (g *GrpcBuilder) Save(events ...adapter.DomainEvent) error {
	grpcEvents := g.grpcEvents(events)
	_, err := g.client.Save(g.ctx, grpcEvents)
	g.deferred()
	return err
}

func (g *GrpcBuilder) Load(uuid uuid.UUID) ([]adapter.DomainEvent, error) {
	id := ac.UUID{Id: uuid.String()}
	var events []adapter.DomainEvent
	response, err := g.client.Load(g.ctx, &id)
	if err == nil {
		events = domainEvents(response.Events.Event)
	}
	g.deferred()
	return events, err
}
func (g *GrpcBuilder) grpcEvents(events []adapter.DomainEvent) *ac.Events {
	grpcEvents := &ac.Events{}
	for _, event := range events {
		grpcEvents.Event = append(grpcEvents.Event, grpcEvent(event))
	}
	return grpcEvents
}

func grpcEvent(event adapter.DomainEvent) *ac.Event {
	return &ac.Event{
		AggregateId: &ac.UUID{Id: event.AggregateId().String()}}
}

func domainEvents(events []*ac.Event) []adapter.DomainEvent {
	var domainEvents []adapter.DomainEvent
	for _, event := range events {
		domainEvents = append(domainEvents, domainEvent(event))
	}
	return domainEvents
}

func domainEvent(event *ac.Event) adapter.DomainEvent {
	id, _ := uuid.Parse(event.AggregateId.Id)
	return StandardEvent{id: id}
}
