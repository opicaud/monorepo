package shape

import (
	"github.com/opicaud/monorepo/events/pkg"
)

func NewShapeCreationCommandHandlerBuilder() *CommandHandlerBuilder {
	return &CommandHandlerBuilder{}
}

func (s *CommandHandlerBuilder) WithEventsFramework(eventsFramework pkg.Provider) *CommandHandlerBuilder {
	s.eventsFramework = eventsFramework
	return s
}

func (s *CommandHandlerBuilder) WithSubscriber(subscriber pkg.Subscriber) *CommandHandlerBuilder {
	s.subscriber = subscriber
	return s
}

func (s *CommandHandlerBuilder) Build() CommandHandler[ShapeCommandApplier] {
	shapeCommandHandler := new(shapeCommandHandler)
	shapeCommandHandler.eventsFramework = s.eventsFramework
	s.eventsFramework.Add(s.subscriber)
	return shapeCommandHandler
}

type CommandHandlerBuilder struct {
	eventsFramework pkg.Provider
	subscriber      pkg.Subscriber
}

type shapeCommandHandler struct {
	eventsFramework pkg.Provider
}

func (f *shapeCommandHandler) Execute(command Command, applier ShapeCommandApplier) error {
	events, err := command.Execute(applier)
	f.eventsFramework.NotifyAll(events...)
	f.eventsFramework.Save(events...)
	return err
}
