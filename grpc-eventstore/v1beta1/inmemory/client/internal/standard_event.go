package eventstore

import "github.com/google/uuid"

// StandardEvent TODO: implement builder

type StandardEvent struct {
	aggregateId uuid.UUID
	name        string
	data        []byte
}

func NewStandardEventForTest(name string) StandardEvent {
	return StandardEvent{aggregateId: uuid.New(), name: name}
}

func NewStandardEvent(aggregateId uuid.UUID, name string, data []byte) StandardEvent {
	return StandardEvent{aggregateId: aggregateId, name: name, data: data}
}

func (s StandardEvent) AggregateId() uuid.UUID {
	return s.aggregateId
}

func (s StandardEvent) Name() string {
	return s.name
}

func (s StandardEvent) Data() []byte {
	return s.data
}
