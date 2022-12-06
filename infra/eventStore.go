package infra

import "github.com/google/uuid"

type EventStore interface {
	Save(events ...Event) error
	Load(id uuid.UUID) []Event
}
