package infra

type StandardEventsEmitter struct {
	subscribers []Subscriber
}

func (s *StandardEventsEmitter) NotifyAll(event ...Event) {
	for _, subscriber := range s.subscribers {
		if subscriber != nil {
			subscriber.Update(event)
		}
	}
}

func (s *StandardEventsEmitter) Add(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}
