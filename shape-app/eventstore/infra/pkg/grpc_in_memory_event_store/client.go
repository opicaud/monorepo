package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/shape-app/eventstore"
	ac "github.com/opicaud/monorepo/shape-app/eventstore/infra/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type StandardEvent struct {
	Id uuid.UUID
}

func (e StandardEvent) AggregateId() uuid.UUID {
	return e.Id
}

type InMemoryGrpcEventStore struct {
	builder GrpcBuilder
}

func NewInMemoryGrpcEventStore() *InMemoryGrpcEventStore {
	return new(InMemoryGrpcEventStore)
}

func (i *InMemoryGrpcEventStore) Save(events ...eventstore.DomainEvent) error {
	err := i.builder.Connect().Save(events...)
	return err
}

func (i *InMemoryGrpcEventStore) Load(id uuid.UUID) ([]eventstore.DomainEvent, error) {
	events, err := i.builder.Connect().Load(id)
	return events, err
}

type GrpcBuilder struct {
	conn   *grpc.ClientConn
	client ac.EventStoreClient
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
	g.client = ac.NewEventStoreClient(g.conn)
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

func (g *GrpcBuilder) Save(events ...eventstore.DomainEvent) error {
	grpcEvents := g.grpcEvents(events)
	_, err := g.client.Save(g.ctx, grpcEvents)
	g.deferred()
	return err
}

func (g *GrpcBuilder) Load(uuid uuid.UUID) ([]eventstore.DomainEvent, error) {
	id := ac.UUID{Id: uuid.String()}
	var events []eventstore.DomainEvent
	response, err := g.client.Load(g.ctx, &id)
	if err == nil {
		events = domainEvents(response.Events.Event)
	}
	g.deferred()
	return events, err
}
func (g *GrpcBuilder) grpcEvents(events []eventstore.DomainEvent) *ac.Events {
	grpcEvents := &ac.Events{}
	for _, event := range events {
		grpcEvents.Event = append(grpcEvents.Event, grpcEvent(event))
	}
	return grpcEvents
}

func grpcEvent(event eventstore.DomainEvent) *ac.Event {
	return &ac.Event{
		AggregateId: &ac.UUID{Id: event.AggregateId().String()}}
}

func domainEvents(events []*ac.Event) []eventstore.DomainEvent {
	var domainEvents []eventstore.DomainEvent
	for _, event := range events {
		domainEvents = append(domainEvents, domainEvent(event))
	}
	return domainEvents
}

func domainEvent(event *ac.Event) eventstore.DomainEvent {
	id, _ := uuid.Parse(event.AggregateId.Id)
	return StandardEvent{Id: id}
}
