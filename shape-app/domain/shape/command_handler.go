package shape

import (
	"github.com/opicaud/monorepo/shape-app/domain/adapter"
)

func NewShapeCreationCommandHandlerBuilder() *CreationCommandHandlerBuilder {
	return &CreationCommandHandlerBuilder{}
}

func (s *CreationCommandHandlerBuilder) WithInfraProvider(infra adapter.Provider) *CreationCommandHandlerBuilder {
	s.provider = infra
	return s
}

func (s *CreationCommandHandlerBuilder) WithSubscriber(subscriber adapter.Subscriber) *CreationCommandHandlerBuilder {
	s.subscriber = subscriber
	return s
}

func (s *CreationCommandHandlerBuilder) Build() CommandHandler {
	shapeCommandHandler := new(shapeCommandHandler)
	shapeCommandHandler.provider = s.provider
	s.provider.Add(s.subscriber)
	return shapeCommandHandler
}

type CreationCommandHandlerBuilder struct {
	provider   adapter.Provider
	subscriber adapter.Subscriber
}

type shapeCommandHandler struct {
	provider adapter.Provider
}

func (f *shapeCommandHandler) Execute(command Command) error {
	events, err := command.Execute(newApplyShapeCommand(f.provider))
	f.provider.NotifyAll(events...)
	f.provider.Save(events...)
	return err
}
