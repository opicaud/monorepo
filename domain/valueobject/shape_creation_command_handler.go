package valueobject

import (
	"example2/domain/commands"
	"fmt"
)

func NewShapeCreationCommandHandler(repository Repository) commands.CommandHandler {
	shapeCommandHandler := new(shapeCommandHandler)
	shapeCommandHandler.repository = repository
	shapeCommandHandler.eventsEmitter = &StandardEventsEmitter{}

	return shapeCommandHandler
}

func newShapeCreationCommandHandlerWithEventsEmitter(repository *InMemoryRepository, emitter EventsEmitter) commands.CommandHandler {
	shapeCommandHandler := new(shapeCommandHandler)
	shapeCommandHandler.repository = repository
	shapeCommandHandler.eventsEmitter = emitter
	return shapeCommandHandler
}

type shapeCommandHandler struct {
	repository    Repository
	eventsEmitter EventsEmitter
}

func (f *shapeCommandHandler) Execute(command commands.Command) error {
	shape, createdEvent := loadShape(command.(newShapeCommand))
	event := applyCommandOnAggregate(command, shape)
	f.eventsEmitter.DispatchEvent(createdEvent, event)
	return f.repository.Save(shape)
}

func loadShape(command newShapeCommand) (Shape, Event) {
	shape, createdEvent, err := newShapeBuilder().createAShape(command.nature).withDimensions(command.dimensions)
	if err != nil {
		panic(err)
	}
	return shape, createdEvent
}

func applyCommandOnAggregate(command commands.Command, shape Shape) Event {
	switch v := command.(type) {
	default:
		fmt.Printf("unexpected command %T", v)
		return nil
	case newShapeCommand:
		return shape.HandleNewShape(command.(newShapeCommand))
	}
}

type EventsEmitter interface {
	DispatchEvent(event ...Event)
}

type StandardEventsEmitter struct{}

func (s *StandardEventsEmitter) DispatchEvent(event ...Event) {}
