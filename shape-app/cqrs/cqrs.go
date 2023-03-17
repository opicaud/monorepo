package cqrs

import "github.com/opicaud/monorepo/events/pkg"

type Command[T interface{}] interface {
	Execute(apply T) ([]pkg.DomainEvent, error)
}

type CommandHandler[K Command[T], T interface{}] interface {
	Execute(command K, commandApplier T) error
}

type CommandHandlerImpl[K Command[T], T any] struct {
	EventsFramework pkg.Provider
}

func (f *CommandHandlerImpl[K, T]) Execute(command K, applier T) error {
	events, err := command.Execute(applier)
	f.EventsFramework.NotifyAll(events...)
	f.EventsFramework.Save(events...)
	return err
}

type CommandHandlerBuilder[T interface{}] struct {
	eventsFramework pkg.Provider
	subscriber      pkg.Subscriber
}

func (s *CommandHandlerBuilder[T]) WithEventsFramework(eventsFramework pkg.Provider) *CommandHandlerBuilder[T] {
	s.eventsFramework = eventsFramework
	return s
}

func (s *CommandHandlerBuilder[T]) WithSubscriber(subscriber pkg.Subscriber) *CommandHandlerBuilder[T] {
	s.subscriber = subscriber
	return s
}

func (s *CommandHandlerBuilder[T]) Build() CommandHandler[Command[T], T] {
	commandHandler := new(CommandHandlerImpl[Command[T], T])
	commandHandler.EventsFramework = s.eventsFramework
	s.eventsFramework.Add(s.subscriber)
	return commandHandler
}
