package infra

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
	infra.EventStore = s.eventStore
	if s.eventsEmitter == nil {
		s.eventsEmitter = &StandardEventsEmitter{}
	}
	infra.EventsEmitter = s.eventsEmitter
	return *infra
}

type Builder struct {
	eventStore    EventStore
	eventsEmitter EventsEmitter
}

type Provider struct {
	EventStore    EventStore
	EventsEmitter EventsEmitter
}
