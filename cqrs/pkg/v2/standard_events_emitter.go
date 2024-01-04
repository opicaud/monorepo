package pkg

import (
	"context"
	"github.com/opicaud/monorepo/events/pkg"
	"sync"
)

type StandardEventsEmitter struct {
	subscribers []Subscriber
}

func (s *StandardEventsEmitter) NotifyAll(ctx context.Context, event ...pkg.DomainEvent) context.Context {
	eventsChn := make(chan []pkg.DomainEvent, len(s.subscribers))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for range s.subscribers {
			eventsChn <- event
		}
	}()
	go func() {
		for _, subscriber := range s.subscribers {
			if subscriber != nil {
				ctx = subscriber.Update(ctx, eventsChn)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	close(eventsChn)
	return ctx
}

func (s *StandardEventsEmitter) Add(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}

type EventsEmitter interface {
	NotifyAll(ctx context.Context, event ...pkg.DomainEvent) context.Context
	Add(subscriber Subscriber)
}

type Subscriber interface {
	Update(ctx context.Context, eventsChn chan []pkg.DomainEvent) context.Context
}
