package valueobject

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHandlerACommand(t *testing.T) {
	inMemoryRepository := NewInMemoryRepository()
	eventsEmitter := MockStandardEventsEmitter{}
	command, _ := newCreationShapeCommand("rectangle", []float32{1, 2})
	handler := newShapeCreationCommandHandlerWithEventsEmitter(inMemoryRepository, &eventsEmitter)

	err := handler.Execute(command.(newShapeCommand))

	expectedEvents := []Event{ShapeCreatedEvent{nature: "rectangle", dimensions: []float32{1, 2}}, AreaShapeCalculated{area: 2}}
	eventsEmitter.mock.AssertCalled(t, "DispatchEvent", expectedEvents)
	assert.NoError(t, err)

}

type MockStandardEventsEmitter struct {
	mock mock.Mock
}

func (s *MockStandardEventsEmitter) DispatchEvent(event ...Event) {
	s.mock.On("DispatchEvent", event)
	s.mock.Called(event)
}

func TestAStandardHandlerACommand(t *testing.T) {
	fakeRepository := NewInMemoryRepository()
	handler := NewShapeCreationCommandHandler(fakeRepository)
	assert.IsType(t, &InMemoryRepository{}, handler.(*shapeCommandHandler).repository)
}
