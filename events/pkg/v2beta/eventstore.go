package pkg

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type EventStore interface {
	pkg.EventStore
	Remove(uuid uuid.UUID) error
	GetHealthClient() grpc_health_v1.HealthClient
}
