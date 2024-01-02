package pkg

import (
	"context"
	pkg2 "github.com/opicaud/monorepo/events/pkg"
	v2beta1 "github.com/opicaud/monorepo/events/pkg/v2beta1"
)

type Command[T interface{}] interface {
	Execute(apply T) ([]pkg2.DomainEvent, error)
}

type CommandHandler[K Command[T], T interface{}] interface {
	Execute(ctx context.Context, command K, commandApplier T) (context.Context, error)
}

type CommandHandlerImpl[K Command[T], T any] struct {
	eventStore    v2beta1.EventStore
	eventsEmitter EventsEmitter
}

func (f *CommandHandlerImpl[K, T]) Execute(ctx context.Context, command K, applier T) (context.Context, error) {
	events, err := command.Execute(applier)
	ctx = f.eventsEmitter.NotifyAll(ctx, events...)
	return ctx, err
}

type CommandHandlerBuilder[T interface{}] struct {
	subscriber           Subscriber
	eventStoreSubscriber EventStoreSubscriber
	eventsEmitter        EventsEmitter
}

func (s *CommandHandlerBuilder[T]) WithSubscriber(subscriber Subscriber) *CommandHandlerBuilder[T] {
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

func (s *CommandHandlerBuilder[T]) WithEventStore(store v2beta1.EventStore) *CommandHandlerBuilder[T] {
	s.eventStoreSubscriber = EventStoreSubscriber{eventStore: store}
	return s
}

func (s *CommandHandlerBuilder[T]) WithEventsEmitter(emitter *StandardEventsEmitter) *CommandHandlerBuilder[T] {
	s.eventsEmitter = emitter
	return s
}

type EventStoreSubscriber struct {
	eventStore v2beta1.EventStore
}

func (e EventStoreSubscriber) Update(ctx context.Context, eventsChn chan []pkg2.DomainEvent) context.Context {
	events := <-eventsChn
	ctx, _, err := e.eventStore.Save(ctx, events...)
	if err != nil {
		panic(err)
	}
	return ctx
}
