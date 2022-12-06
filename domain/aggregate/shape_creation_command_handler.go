package aggregate

import (
	"example2/infra"
)

func NewShapeCreationCommandHandlerBuilder() *ShapeCreationCommandHandlerBuilder {
	return &ShapeCreationCommandHandlerBuilder{}
}

func (s *ShapeCreationCommandHandlerBuilder) WithEventStore(eventStore infra.EventStore) *ShapeCreationCommandHandlerBuilder {
	s.eventStore = eventStore
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
	if s.eventStore == nil {
		s.eventStore = infra.NewInMemoryEventStore()
	}
	shapeCommandHandler.eventstore = s.eventStore
	shapeCommandHandler.subscriber = s.subscriber
	if s.eventsEmitter == nil {
		s.eventsEmitter = &infra.StandardEventsEmitter{}
	}
	shapeCommandHandler.eventsEmitter = s.eventsEmitter
	shapeCommandHandler.eventsEmitter.Add(s.subscriber)
	return shapeCommandHandler
}

type ShapeCreationCommandHandlerBuilder struct {
	eventStore    infra.EventStore
	eventsEmitter infra.EventsEmitter
	subscriber    infra.Subscriber
}

type shapeCommandHandler struct {
	eventstore    infra.EventStore
	eventsEmitter infra.EventsEmitter
	subscriber    infra.Subscriber
}

func (f *shapeCommandHandler) Execute(command ShapeCommand) error {
	events := command.Apply(newApplyShapeCommand(f.eventstore))
	f.eventsEmitter.NotifyAll(events...)
	return f.eventstore.Save(events...)
}
