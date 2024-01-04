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
	eventStore    pkg.EventStore
	eventsEmitter pkg.EventsEmitter
}

func (f *CommandHandlerImpl[K, T]) Execute(command K, applier T) error {
	events, err := command.Execute(applier)
	f.eventsEmitter.NotifyAll(events...)
	return err
}

type CommandHandlerBuilder[T interface{}] struct {
	subscriber           pkg.Subscriber
	eventStoreSubscriber EventStoreSubscriber
	eventsEmitter        pkg.EventsEmitter
}

func (s *CommandHandlerBuilder[T]) WithSubscriber(subscriber pkg.Subscriber) *CommandHandlerBuilder[T] {
	s.subscriber = subscriber
	return s
}

func (s *CommandHandlerBuilder[T]) Build() CommandHandler[Command[T], T] {
	commandHandler := new(CommandHandlerImpl[Command[T], T])
	commandHandler.eventsEmitter = s.eventsEmitter
	s.eventsEmitter.Add(s.subscriber)
	s.eventsEmitter.Add(s.eventStoreSubscriber)
	return commandHandler
}

func (s *CommandHandlerBuilder[T]) WithEventStore(store pkg.EventStore) *CommandHandlerBuilder[T] {
	s.eventStoreSubscriber = EventStoreSubscriber{eventStore: store}
	return s
}

func (s *CommandHandlerBuilder[T]) WithEventsEmitter(emitter pkg.EventsEmitter) *CommandHandlerBuilder[T] {
	s.eventsEmitter = emitter
	return s
}

type EventStoreSubscriber struct {
	eventStore pkg.EventStore
}

func (e EventStoreSubscriber) Update(events []pkg.DomainEvent) {
	err := e.eventStore.Save(events...)
	if err != nil {
		panic(err)
	}
}
