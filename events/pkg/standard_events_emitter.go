package pkg

type StandardEventsEmitter struct {
	subscribers []Subscriber
}

func (s *StandardEventsEmitter) NotifyAll(event ...DomainEvent) {
	for _, subscriber := range s.subscribers {
		if subscriber != nil {
			subscriber.Update(event)
		}
	}
}

func (s *StandardEventsEmitter) Add(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}

type EventsEmitter interface {
	NotifyAll(event ...DomainEvent)
	Add(subscriber Subscriber)
}

type Subscriber interface {
	Update(events []DomainEvent)
}
