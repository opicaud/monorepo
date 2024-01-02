package pkg

import (
	"context"
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type EventStore interface {
	Save(ctx context.Context, events ...pkg.DomainEvent) (context.Context, []pkg.DomainEvent, error)
	Load(ctx context.Context, id uuid.UUID) (context.Context, []pkg.DomainEvent, error)
	Remove(ctx context.Context, uuid uuid.UUID) (context.Context, error)
	GetHealthClient(ctx context.Context) (context.Context, grpc_health_v1.HealthClient)
}
