package infra

import "github.com/google/uuid"

func NewInfraBuilder() *Builder {
	return &Builder{}
}

func (s *Builder) WithEventStore(eventStore EventStore) *Builder {
	s.eventStore = eventStore
	return s
}

func (s *Builder) WithEmitter(emitter EventsEmitter) *Builder {
	s.eventsEmitter = emitter
	return s
}

func (s *Builder) Build() Provider {
	infra := new(Provider)
	if s.eventStore == nil {
		s.eventStore = NewInMemoryEventStore()
	}
	infra.eventstore = s.eventStore
	if s.eventsEmitter == nil {
		s.eventsEmitter = &StandardEventsEmitter{}
	}
	infra.eventsEmitter = s.eventsEmitter
	return *infra
}

type Builder struct {
	eventStore    EventStore
	eventsEmitter EventsEmitter
}

type Provider struct {
	eventstore    EventStore
	eventsEmitter EventsEmitter
}

func (f *Provider) NotifyAll(event ...Event) {
	f.eventsEmitter.NotifyAll(event...)
}

func (f *Provider) Add(subscriber Subscriber) {
	f.eventsEmitter.Add(subscriber)
}

func (f *Provider) Save(events ...Event) {
	f.eventstore.Save(events...)
}

func (f *Provider) Load(uuid uuid.UUID) ([]Event, error) {
	return f.eventstore.Load(uuid)
}
