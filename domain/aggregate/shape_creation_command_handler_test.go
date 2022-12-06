package aggregate

import (
	"example2/infra"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	inMemoryRepository := NewInMemoryEventStore()
	eventsEmitter := infra.StandardEventsEmitter{}
	subscriber := SubscriberForTest{}
	command, _ := newCreationShapeCommand("rectangle", []float32{1, 2})
	handler := NewShapeCreationCommandHandlerBuilder().
		WithEventStore(inMemoryRepository).
		WithEmitter(&eventsEmitter).
		WithSubscriber(&subscriber).
		Build()

	err := handler.Execute(command)

	assert.Equal(t, 2, len(subscriber.events))
	assert.Equal(t, subscriber.ids[0], subscriber.ids[1])

	assert.Equal(t, ShapeCreatedEvent{id: subscriber.ids[0], Nature: "rectangle", dimensions: []float32{1, 2}}, subscriber.events[0])
	assert.Equal(t, AreaShapeCalculated{id: subscriber.ids[1], Area: 2}, subscriber.events[1])
	assert.NoError(t, err)

}

type SubscriberForTest struct {
	events []infra.Event
	ids    []uuid.UUID
}

func (s *SubscriberForTest) Update(events []infra.Event) {
	s.events = events
	for _, e := range s.events {
		s.ids = append(s.ids, e.AggregateId())
	}
}

func TestAStandardHandlerACommand(t *testing.T) {
	handler := NewShapeCreationCommandHandlerBuilder().
		WithEventStore(NewInMemoryEventStore()).
		Build()
	assert.IsType(t, &InMemoryRepository{}, handler.(*shapeCommandHandler).eventstore)
}
