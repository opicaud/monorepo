package infra

import "github.com/google/uuid"

type EventsEmitter interface {
	NotifyAll(event ...Event)
	Add(subscriber Subscriber)
}

type StandardEventsEmitter struct {
	subscribers []Subscriber
}

type Event interface {
	AggregateId() uuid.UUID
}

type Subscriber interface {
	Update(events []Event)
}

func (s *StandardEventsEmitter) NotifyAll(event ...Event) {
	for _, s := range s.subscribers {
		s.Update(event)
	}
}

func (s *StandardEventsEmitter) Add(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}
