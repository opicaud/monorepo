package pkg

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	v2beta "github.com/opicaud/monorepo/events/pkg/v2beta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		eventStore := &FakeEventStore{}
		eventsEmitter := &v2beta.StandardEventsEmitter{}
		eventStore.mock.On("Save", nil).Return()
		subscriber := &FakeSubscriber{}
		subscriber.mock.On("Update", nil).Return()
		commandHandler := c.WithEventStore(eventStore).
			WithEventsEmitter(eventsEmitter).
			WithSubscriber(subscriber).
			Build()
		err := commandHandler.Execute(f, v)

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

func (f *FakeEventStore) Save(events ...pkg.DomainEvent) error {
	f.mock.Called(nil)
	return nil
}

func (f *FakeEventStore) Load(id uuid.UUID) ([]pkg.DomainEvent, error) {
	panic("implement me")
}

type FakeSubscriber struct {
	mock mock.Mock
}

func (g *FakeSubscriber) Update(eventsChn chan []pkg.DomainEvent) {
	g.mock.Called(nil)
}
