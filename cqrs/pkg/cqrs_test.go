package pkg

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
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

	t.Run("v1", v1(c, subscriber, f, v))
	t.Run("v2", v2(c, subscriber, f, v))

}

func v2(c CommandHandlerBuilder[FakeCommandApplier], subscriber *FakeSubscriber, f FakeCommand[FakeCommandApplier], v FakeCommandApplier) func(t *testing.T) {
	return func(t *testing.T) {
		eventStore := &FakeEventStore{}
		store := &FakeEventStore{}
		eventStore.mock.On("Save", nil).Return()
		eventStoreFramework := pkg.NewEventsFrameworkBuilder().WithEventStore(store).Build()
		commandHandler := c.WithEventStore(eventStore).
			WithEventsFramework(eventStoreFramework).
			WithSubscriber(subscriber).
			Build()
		err := commandHandler.Execute(f, v)

		assert.NoError(t, err)

		eventStore.mock.AssertCalled(t, "Save", nil)
		subscriber.mock.AssertCalled(t, "Update", nil)
		store.mock.AssertNotCalled(t, "Save", nil)

	}
}

func v1(c CommandHandlerBuilder[FakeCommandApplier], subscriber *FakeSubscriber, f FakeCommand[FakeCommandApplier], v FakeCommandApplier) func(t *testing.T) {
	return func(t *testing.T) {
		store := &FakeEventStore{}
		store.mock.On("Save", nil).Return()
		eventStoreFramework := pkg.NewEventsFrameworkBuilder().WithEventStore(store).Build()
		commandHandler := c.WithEventsFramework(eventStoreFramework).WithSubscriber(subscriber).Build()
		err := commandHandler.Execute(f, v)

		assert.NoError(t, err)
		store.mock.AssertCalled(t, "Save", nil)
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

func (f *FakeSubscriber) Update(events []pkg.DomainEvent) {
	f.mock.Called(nil)
}
