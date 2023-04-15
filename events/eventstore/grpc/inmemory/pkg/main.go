package pkg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/eventstore/grpc/inmemory/internal"
	gen "github.com/opicaud/monorepo/events/eventstore/grpc/proto"
	"github.com/opicaud/monorepo/events/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type InMemoryGrpcEventStore struct {
	builder GrpcBuilder
}

func (i *InMemoryGrpcEventStore) Remove(uuid uuid.UUID) error {
	return i.builder.Connect().Remove(uuid)
}

func NewInMemoryGrpcEventStoreFrom(address string, port int) *InMemoryGrpcEventStore {
	return newInMemoryGrpcEventStore(address, port)
}

func newInMemoryGrpcEventStore(address string, port int) *InMemoryGrpcEventStore {
	i := new(InMemoryGrpcEventStore)
	i.builder.setAddress(address)
	i.builder.setPort(port)
	return i
}

func NewInMemoryGrpcEventStore() *InMemoryGrpcEventStore {
	return newInMemoryGrpcEventStore("localhost", 50051)
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
	conn    *grpc.ClientConn
	client  gen.EventStoreClient
	ctx     context.Context
	cancel  context.CancelFunc
	address string
	err     error
	port    int
}

func (g *GrpcBuilder) Connect() *GrpcBuilder {
	g.conn, g.err = grpc.Dial(fmt.Sprintf("%s:%d", g.address, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func (g *GrpcBuilder) Remove(uuid uuid.UUID) error {
	id := gen.UUID{Id: uuid.String()}
	_, err := g.client.Remove(g.ctx, &id)
	g.deferred()
	return err
}

func (g *GrpcBuilder) Load(uuid uuid.UUID) ([]pkg.DomainEvent, error) {
	id := gen.UUID{Id: uuid.String()}
	log.Printf("did to search: %s\n", id.Id)
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

func (g *GrpcBuilder) setAddress(address string) {
	g.address = address

}

func (g *GrpcBuilder) setPort(port int) {
	g.port = port

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
