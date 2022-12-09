package shape

import (
	"example2/infra"
)

func NewShapeCreationCommandHandlerBuilder() *ShapeCreationCommandHandlerBuilder {
	return &ShapeCreationCommandHandlerBuilder{}
}

func (s *ShapeCreationCommandHandlerBuilder) WithInfraProvider(infra infra.Provider) *ShapeCreationCommandHandlerBuilder {
	s.provider = infra
	return s
}

func (s *ShapeCreationCommandHandlerBuilder) WithSubscriber(subscriber infra.Subscriber) *ShapeCreationCommandHandlerBuilder {
	s.subscriber = subscriber
	return s
}

func (s *ShapeCreationCommandHandlerBuilder) Build() ShapeCommandHandler {
	shapeCommandHandler := new(shapeCommandHandler)
	shapeCommandHandler.provider = s.provider
	s.provider.Add(s.subscriber)
	return shapeCommandHandler
}

type ShapeCreationCommandHandlerBuilder struct {
	provider   infra.Provider
	subscriber infra.Subscriber
}

type shapeCommandHandler struct {
	provider infra.Provider
}

func (f *shapeCommandHandler) Execute(command ShapeCommand) error {
	events := command.Apply(newApplyShapeCommand(f.provider))
	f.provider.NotifyAll(events...)
	return f.provider.Save(events...)
}
