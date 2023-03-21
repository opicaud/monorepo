package pkg

import (
	"fmt"
	"github.com/google/uuid"
)

// NewEventsFrameworkBuilder Deprecated
func NewEventsFrameworkBuilder() *Builder {
	return &Builder{}
}

// WithEventStore NewEventsFrameworkBuilder Deprecated
func (s *Builder) WithEventStore(eventStore EventStore) *Builder {
	s.eventStore = eventStore
	return s
}

// WithEventsEmitter Deprecated
func (s *Builder) WithEventsEmitter(eventsEmitter EventsEmitter) *Builder {
	s.eventsEmitter = eventsEmitter
	return s
}

// Build Deprecated
func (s *Builder) Build() Provider {
	infra := new(Provider)
	if s.eventStore == nil {
		panic("eventStore impl must be provided")
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

// Provider Deprecated
type Provider struct {
	eventstore    EventStore
	eventsEmitter EventsEmitter
}

func (f *Provider) NotifyAll(event ...DomainEvent) {
	f.eventsEmitter.NotifyAll(event...)
}

func (f *Provider) Add(subscriber Subscriber) {
	f.eventsEmitter.Add(subscriber)
}

func (f *Provider) Save(events ...DomainEvent) {
	err := f.eventstore.Save(events...)
	if err != nil {
		err = fmt.Errorf("error has occured when save events")
		fmt.Println(err.Error())
	}
}

func (f *Provider) Load(uuid uuid.UUID) ([]DomainEvent, error) {
	return f.eventstore.Load(uuid)
}
