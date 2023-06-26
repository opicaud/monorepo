package pkg

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
)

type EventStore interface {
	pkg.EventStore
	Remove(uuid uuid.UUID) error
}
