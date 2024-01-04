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
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
)

func NewInMemoryGrpcEventStoreFrom(address string, port int) *InMemoryGrpcEventStore {
	return newInMemoryGrpcEventStore(address, port)
}

func newInMemoryGrpcEventStore(address string, port int) *InMemoryGrpcEventStore {
	i := new(InMemoryGrpcEventStore)
	i.port = port
	i.address = address
	i.Connect()
	return i
}

type InMemoryGrpcEventStore struct {
	conn            *grpc.ClientConn
	client          gen.EventStoreClient
	cancel          context.CancelFunc
	address         string
	err             error
	port            int
	healthIndicator grpc_health_v1.HealthClient
}

func (g *InMemoryGrpcEventStore) Connect() *InMemoryGrpcEventStore {
	g.conn, g.err = grpc.Dial(fmt.Sprintf("%s:%d", g.address, g.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if g.err != nil {
		log.Panic(g.err)
	}

	g.client = gen.NewEventStoreClient(g.conn)
	g.healthIndicator = grpc_health_v1.NewHealthClient(g.conn)
	return g
}

func (g *InMemoryGrpcEventStore) deferred() {
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}(g.conn)
	defer g.cancel()
}

func (g *InMemoryGrpcEventStore) GetHealthClient(ctx context.Context) (context.Context, grpc_health_v1.HealthClient) {
	return ctx, g.healthIndicator
}

func (g *InMemoryGrpcEventStore) Save(ctx context.Context, events ...pkg.DomainEvent) (context.Context, []pkg.DomainEvent, error) {
	grpcEvents := g.grpcEvents(events)
	_, err := g.client.Save(ctx, grpcEvents)
	g.deferred()
	return ctx, events, err
}

func (g *InMemoryGrpcEventStore) Remove(ctx context.Context, uuid uuid.UUID) (context.Context, error) {
	id := gen.UUID{Id: uuid.String()}
	_, err := g.client.Remove(ctx, &id)
	g.deferred()
	return ctx, err
}

func (g *InMemoryGrpcEventStore) Load(ctx context.Context, uuid uuid.UUID) (context.Context, []pkg.DomainEvent, error) {
	id := gen.UUID{Id: uuid.String()}
	log.Printf("did to search: %s\n", id.Id)
	var events []pkg.DomainEvent
	response, err := g.client.Load(ctx, &id)
	if err == nil {
		events = domainEvents(response.Events.Event)
	}
	g.deferred()
	return ctx, events, err
}
func (g *InMemoryGrpcEventStore) grpcEvents(events []pkg.DomainEvent) *gen.Events {
	grpcEvents := &gen.Events{}
	for _, event := range events {
		grpcEvents.Event = append(grpcEvents.Event, grpcEvent(event))
	}
	return grpcEvents
}

func (g *InMemoryGrpcEventStore) setAddress(address string) {
	g.address = address

}

func (g *InMemoryGrpcEventStore) setPort(port int) {
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
