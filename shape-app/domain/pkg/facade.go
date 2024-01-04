package pkg

import (
	"github.com/google/uuid"
	cqrs "github.com/opicaud/monorepo/cqrs/pkg/v3beta1"
	"github.com/opicaud/monorepo/shape-app/domain/internal"
)

func (f *shapeFacade) NewCreationShapeCommand(nature string, dimensions ...float32) cqrs.Command[internal.CommandApplier] {
	command := internal.NewCreationShapeCommand(nature, dimensions)
	return command
}

func (f *shapeFacade) NewStretchShapeCommand(id uuid.UUID, stretchBy float32) cqrs.Command[internal.CommandApplier] {
	command := internal.NewStretchShapeCommand(id, stretchBy)
	return command
}

func New() ShapeFacade {
	return new(shapeFacade)
}

type shapeFacade struct{}

type ShapeFacade interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) cqrs.Command[internal.CommandApplier]
	NewStretchShapeCommand(id uuid.UUID, stretchBy float32) cqrs.Command[internal.CommandApplier]
	NewCommandHandlerBuilder() *cqrs.CommandHandlerBuilder[internal.CommandApplier]
	NewShapeCommandApplier(eventsFramework cqrs.EventStore) internal.CommandApplier
}

func (f *shapeFacade) NewCommandHandlerBuilder() *cqrs.CommandHandlerBuilder[internal.CommandApplier] {
	return &cqrs.CommandHandlerBuilder[internal.CommandApplier]{}
}

func (f *shapeFacade) NewShapeCommandApplier(eventStore cqrs.EventStore) internal.CommandApplier {
	return internal.NewShapeCommandApplier(eventStore)
}
