package pkg

import (
	"github.com/google/uuid"
)

type EventStore interface {
	Save(events ...DomainEvent) error
	Load(id uuid.UUID) ([]DomainEvent, error)
}
