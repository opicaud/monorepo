package aggregate

import (
	"example2/domain/commands"
	"example2/infra"
	"fmt"
)

func NewShapeCreationCommandHandlerBuilder() *ShapeCreationCommandHandlerBuilder {
	return &ShapeCreationCommandHandlerBuilder{
		eventsEmitter: &infra.StandardEventsEmitter{},
		repository:    NewInMemoryRepository(),
	}
}

func (s *ShapeCreationCommandHandlerBuilder) WithRepository(repository Repository) *ShapeCreationCommandHandlerBuilder {
	s.repository = repository
	return s
}

func (s *ShapeCreationCommandHandlerBuilder) WithEmitter(emitter infra.EventsEmitter) *ShapeCreationCommandHandlerBuilder {
	s.eventsEmitter = emitter
	return s
}

func (s *ShapeCreationCommandHandlerBuilder) WithSubscriber(subscriber infra.Subscriber) *ShapeCreationCommandHandlerBuilder {
	s.subscriber = subscriber
	return s
}

func (s *ShapeCreationCommandHandlerBuilder) Build() commands.CommandHandler {
	shapeCommandHandler := new(shapeCommandHandler)
	shapeCommandHandler.repository = s.repository
	shapeCommandHandler.eventsEmitter = s.eventsEmitter
	shapeCommandHandler.eventsEmitter.Add(s.subscriber)
	return shapeCommandHandler
}

type ShapeCreationCommandHandlerBuilder struct {
	repository    Repository
	eventsEmitter infra.EventsEmitter
	subscriber    infra.Subscriber
}

type shapeCommandHandler struct {
	repository    Repository
	eventsEmitter infra.EventsEmitter
	subscriber    infra.Subscriber
}

func (f *shapeCommandHandler) Execute(command commands.Command) error {
	shape, createdEvent := loadShape(command.(newShapeCommand))
	event := applyCommandOnAggregate(command, shape)
	f.eventsEmitter.NotifyAll(createdEvent, event)
	return f.repository.Save(shape)
}

func loadShape(command newShapeCommand) (Shape, infra.Event) {
	shape, createdEvent, err := newShapeBuilder().createAShape(command.nature).withDimensions(command.dimensions)
	if err != nil {
		panic(err)
	}
	return shape, createdEvent
}

func applyCommandOnAggregate(command commands.Command, shape Shape) infra.Event {
	switch v := command.(type) {
	default:
		fmt.Printf("unexpected command %T", v)
		return nil
	case newShapeCommand:
		return shape.HandleNewShape(command.(newShapeCommand))
	}
}
