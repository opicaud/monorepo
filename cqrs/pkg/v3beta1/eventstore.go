package pkg

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type EventStore interface {
	Save(ctx context.Context, events ...DomainEvent) (context.Context, []DomainEvent, error)
	Load(ctx context.Context, id uuid.UUID) (context.Context, []DomainEvent, error)
	Remove(ctx context.Context, uuid uuid.UUID) (context.Context, error)
	GetHealthClient(ctx context.Context) (context.Context, grpc_health_v1.HealthClient)
}
