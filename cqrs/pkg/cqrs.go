package pkg

import (
	"github.com/opicaud/monorepo/events/pkg"
)

type Command[T interface{}] interface {
	Execute(apply T) ([]pkg.DomainEvent, error)
}

type CommandHandler[K Command[T], T interface{}] interface {
	Execute(command K, commandApplier T) error
}

type CommandHandlerImpl[K Command[T], T any] struct {
	EventsFramework pkg.Provider
	eventStore      pkg.EventStore
	eventsEmitter   pkg.EventsEmitter
}

func (f *CommandHandlerImpl[K, T]) Execute(command K, applier T) error {
	events, err := command.Execute(applier)
	if f.eventStore != nil {
		_ = f.eventStore.Save(events...)
	} else {
		f.EventsFramework.Save(events...)
	}
	if f.eventsEmitter != nil {
		f.eventsEmitter.NotifyAll(events...)
	} else {
		f.EventsFramework.NotifyAll(events...)
	}

	return err
}

type CommandHandlerBuilder[T interface{}] struct {
	eventsFramework pkg.Provider
	subscriber      pkg.Subscriber
	eventStore      pkg.EventStore
	eventsEmitter   pkg.EventsEmitter
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
	commandHandler.eventStore = s.eventStore
	commandHandler.eventsEmitter = s.eventsEmitter
	if s.eventsEmitter == nil {
		s.eventsFramework.Add(s.subscriber)
	} else {
		s.eventsEmitter.Add(s.subscriber)
	}

	return commandHandler
}

func (s *CommandHandlerBuilder[T]) WithEventStore(store pkg.EventStore) *CommandHandlerBuilder[T] {
	s.eventStore = store
	return s
}

func (s *CommandHandlerBuilder[T]) WithEventsEmitter(emitter pkg.EventsEmitter) *CommandHandlerBuilder[T] {
	s.eventsEmitter = emitter
	return s
}
