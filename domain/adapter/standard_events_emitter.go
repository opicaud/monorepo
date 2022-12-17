package adapter

type standardEventsEmitter struct {
	subscribers []Subscriber
}

func (s *standardEventsEmitter) NotifyAll(event ...DomainEvent) {
	for _, subscriber := range s.subscribers {
		if subscriber != nil {
			subscriber.Update(event)
		}
	}
}

func (s *standardEventsEmitter) Add(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}
