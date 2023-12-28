package pkg

import (
	"github.com/opicaud/monorepo/events/pkg"
	"sync"
)

type StandardEventsEmitter struct {
	subscribers []Subscriber
}

func (s *StandardEventsEmitter) NotifyAll(event ...pkg.DomainEvent) {
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
				subscriber.Update(eventsChn)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	close(eventsChn)
}

func (s *StandardEventsEmitter) Add(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}

type EventsEmitter interface {
	NotifyAll(event ...pkg.DomainEvent)
	Add(subscriber Subscriber)
}

type Subscriber interface {
	Update(eventsChn chan []pkg.DomainEvent)
}
