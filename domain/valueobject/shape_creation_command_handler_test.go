package valueobject

import (
	"example2/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	inMemoryRepository := NewInMemoryRepository()
	eventsEmitter := infra.StandardEventsEmitter{}
	subscriber := SubscriberForTest{}
	command, _ := newCreationShapeCommand("rectangle", []float32{1, 2})
	handler := NewShapeCreationCommandHandlerBuilder().
		WithRepository(inMemoryRepository).
		WithEmitter(&eventsEmitter).
		WithSubscriber(&subscriber).
		Build()

	err := handler.Execute(command)
	assert.Equal(t, 2, len(subscriber.events))
	assert.Equal(t, ShapeCreatedEvent{nature: "rectangle", dimensions: []float32{1, 2}}, subscriber.events[0])
	assert.Equal(t, AreaShapeCalculated{Area: 2}, subscriber.events[1])
	assert.NoError(t, err)

}

type SubscriberForTest struct {
	events []infra.Event
}

func (s *SubscriberForTest) Update(events []infra.Event) {
	s.events = events
}

func TestAStandardHandlerACommand(t *testing.T) {
	handler := NewShapeCreationCommandHandlerBuilder().
		WithRepository(NewInMemoryRepository()).
		Build()
	assert.IsType(t, &InMemoryRepository{}, handler.(*shapeCommandHandler).repository)
}
