package test

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	cqrs "github.com/opicaud/monorepo/cqrs/v3/pkg"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewFakeInMemoryEventStore() *InMemoryEventStore {
	fakeRepository := new(InMemoryEventStore)
	fakeRepository.MockedHealthClient = MockedHealthClient{Mock: new(mock.Mock)}
	return fakeRepository
}

func (f *InMemoryEventStore) Save(events ...cqrs.DomainEvent) error {
	f.events = append(f.events, events...)
	return nil
}

func (f *InMemoryEventStore) Load(uuid uuid.UUID) ([]cqrs.DomainEvent, error) {
	w := 0
	for _, e := range f.events {
		if e.AggregateId() == uuid {
			f.events[w] = e
			w++
		}
	}
	if len(f.events[0:w]) == 0 {
		return nil, fmt.Errorf("No aggregate with Id %s has been found", uuid)
	}
	return f.events[0:w], nil
}

func (f *InMemoryEventStore) Remove(uuid uuid.UUID) error {
	panic("not implemented")
}

type InMemoryEventStore struct {
	events             []cqrs.DomainEvent
	MockedHealthClient MockedHealthClient
}

func (f *InMemoryEventStore) GetHealthClient() grpc_health_v1.HealthClient {
	return &f.MockedHealthClient
}

type MockedHealthClient struct {
	Mock *mock.Mock
}

func (m *MockedHealthClient) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest, opts ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error) {
	args := m.Mock.Called()
	response := args.Get(0).(*grpc_health_v1.HealthCheckResponse)
	return response, args.Error(1)
}

func (m *MockedHealthClient) Watch(ctx context.Context, in *grpc_health_v1.HealthCheckRequest, opts ...grpc.CallOption) (grpc_health_v1.Health_WatchClient, error) {
	panic("implement me")
}
