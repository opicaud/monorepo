package shape

import "github.com/opicaud/monorepo/shape-app/events/pkg"

func NewShapeCreationCommandHandlerBuilder() *CreationCommandHandlerBuilder {
	return &CreationCommandHandlerBuilder{}
}

func (s *CreationCommandHandlerBuilder) WithInfraProvider(infra pkg.Provider) *CreationCommandHandlerBuilder {
	s.provider = infra
	return s
}

func (s *CreationCommandHandlerBuilder) WithSubscriber(subscriber pkg.Subscriber) *CreationCommandHandlerBuilder {
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
	provider   pkg.Provider
	subscriber pkg.Subscriber
}

type shapeCommandHandler struct {
	provider pkg.Provider
}

func (f *shapeCommandHandler) Execute(command Command) error {
	events, err := command.Execute(newApplyShapeCommand(f.provider))
	f.provider.NotifyAll(events...)
	f.provider.Save(events...)
	return err
}
