package pkg

import "github.com/google/uuid"

type DomainEvent interface {
	AggregateId() uuid.UUID
	Name() string
	Data() []byte
}
