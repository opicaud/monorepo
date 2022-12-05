package aggregate

import (
	"example2/infra"
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

func (s *ShapeCreationCommandHandlerBuilder) Build() ShapeCommandHandler {
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

func (f *shapeCommandHandler) Execute(command ShapeCommand) error {
	shape, events := command.Apply(newApplyShapeCommand())
	f.eventsEmitter.NotifyAll(events...)
	return f.repository.Save(shape)
}

type ApplyShapeCommandImpl struct{}

func newApplyShapeCommand() ApplyShapeCommand {
	return new(ApplyShapeCommandImpl)
}

func (ApplyShapeCommandImpl) ApplyNewShapeCommand(command newShapeCommand) (Shape, []infra.Event) {
	shape, shapeCreatedEvent, err := newShapeBuilder().createAShape(command.nature).withDimensions(command.dimensions)
	if err != nil {
		panic(err)
	}
	areaShapeCalculated := shape.HandleCaculateShapeArea(command)
	return shape, []infra.Event{shapeCreatedEvent, areaShapeCalculated}
}

func (ApplyShapeCommandImpl) ApplyNewStretchCommand(command newStretchCommand) (Shape, []infra.Event) {
	return nil, nil
}
