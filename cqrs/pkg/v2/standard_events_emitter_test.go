package pkg

import (
	"context"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldUseStandardEmitter(t *testing.T) {
	ctx := context.Background()
	emitter := StandardEventsEmitter{}
	subscriber := SubscriberForTest{}
	emitter.Add(&subscriber)
	emitter.Add(&SubscriberForTest{})
	ctx = emitter.NotifyAll(ctx, StandardEvent{aggregateId: uuid.New(), name: "test"})

	assert.Len(t, subscriber.eventsFromUpdate, 1)
	assert.Equal(t, subscriber.eventsFromUpdate[0].Name(), "test")

}

type SubscriberForTest struct {
	eventsFromUpdate []pkg.DomainEvent
}

func (s *SubscriberForTest) Update(ctx context.Context, eventsChn chan []pkg.DomainEvent) context.Context {
	s.eventsFromUpdate = <-eventsChn
	return ctx
}

type StandardEvent struct {
	aggregateId uuid.UUID
	name        string
	data        []byte
}

func (s StandardEvent) AggregateId() uuid.UUID {
	return s.aggregateId
}

func (s StandardEvent) Name() string {
	return s.name
}

func (s StandardEvent) Data() []byte {
	return nil
}
