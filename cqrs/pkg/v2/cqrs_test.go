package pkg

import (
	"context"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/health/grpc_health_v1"
	"testing"
)

func TestExecuteACommand(t *testing.T) {
	subscriber := &FakeSubscriber{}
	subscriber.mock.On("Update", nil).Return()
	c := CommandHandlerBuilder[FakeCommandApplier]{}
	f := FakeCommand[FakeCommandApplier]{}
	v := FakeCommandApplier{}

	t.Run("v2", v2(c, f, v))

}

func v2(c CommandHandlerBuilder[FakeCommandApplier], f FakeCommand[FakeCommandApplier], v FakeCommandApplier) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		eventStore := &FakeEventStore{}
		eventsEmitter := &StandardEventsEmitter{}
		eventStore.mock.On("Save", nil).Return()
		subscriber := &FakeSubscriber{}
		subscriber.mock.On("Update", nil).Return()
		commandHandler := c.WithEventStore(eventStore).
			WithEventsEmitter(eventsEmitter).
			WithSubscriber(subscriber).
			Build()
		ctx, err := commandHandler.Execute(ctx, f, v)

		assert.NoError(t, err)

		eventStore.mock.AssertCalled(t, "Save", nil)
		subscriber.mock.AssertCalled(t, "Update", nil)

	}
}

type FakeCommandApplier struct{}

type FakeCommand[T FakeCommandApplier] struct{}

func (f FakeCommand[T]) Execute(apply T) ([]pkg.DomainEvent, error) {
	return nil, nil
}

type FakeEventStore struct {
	mock mock.Mock
}

func (f *FakeEventStore) Save(ctx context.Context, events ...pkg.DomainEvent) (context.Context, []pkg.DomainEvent, error) {
	f.mock.Called(nil)
	return ctx, events, nil
}

func (f *FakeEventStore) Load(ctx context.Context, id uuid.UUID) (context.Context, []pkg.DomainEvent, error) {
	panic("implement me")
}

func (f *FakeEventStore) Remove(ctx context.Context, uuid uuid.UUID) (context.Context, error) {
	return ctx, nil
}
func (f *FakeEventStore) GetHealthClient(ctx context.Context) (context.Context, grpc_health_v1.HealthClient) {
	return ctx, nil
}

type FakeSubscriber struct {
	mock mock.Mock
}

func (g *FakeSubscriber) Update(ctx context.Context, eventsChn chan []pkg.DomainEvent) context.Context {
	g.mock.Called(nil)
	return ctx
}
