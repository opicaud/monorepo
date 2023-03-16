package shape

import (
	"github.com/opicaud/monorepo/events/pkg"
)

func NewCommandHandlerBuilder() *CommandHandlerBuilder {
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

func (s *CommandHandlerBuilder) Build() CommandHandler[Command[CommandApplier], CommandApplier] {
	commandHandler := new(CommandHandlerImpl[Command[CommandApplier], CommandApplier])
	commandHandler.eventsFramework = s.eventsFramework
	s.eventsFramework.Add(s.subscriber)
	return commandHandler
}

type CommandHandlerBuilder struct {
	eventsFramework pkg.Provider
	subscriber      pkg.Subscriber
}

type CommandHandlerImpl[K Command[T], T any] struct {
	eventsFramework pkg.Provider
}

func (f *CommandHandlerImpl[K, T]) Execute(command K, applier T) error {
	events, err := command.Execute(applier)
	f.eventsFramework.NotifyAll(events...)
	f.eventsFramework.Save(events...)
	return err
}
