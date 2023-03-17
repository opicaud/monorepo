package pkg

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/opicaud/monorepo/shape-app/cqrs"
	"github.com/opicaud/monorepo/shape-app/domain/shape/internal"
)

func (f *shapeFacade) NewCreationShapeCommand(nature string, dimensions ...float32) internal.Command[internal.CommandApplier] {
	command := internal.NewCreationShapeCommand(nature, dimensions)
	return command
}

func (f *shapeFacade) NewStretchShapeCommand(id uuid.UUID, stretchBy float32) internal.Command[internal.CommandApplier] {
	command := internal.NewStretchShapeCommand(id, stretchBy)
	return command
}

func New() ShapeFacade {
	return new(shapeFacade)
}

type shapeFacade struct{}

type ShapeFacade interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) internal.Command[internal.CommandApplier]
	NewStretchShapeCommand(id uuid.UUID, stretchBy float32) internal.Command[internal.CommandApplier]
	NewCommandHandlerBuilder() *cqrs.CommandHandlerBuilder[internal.CommandApplier]
	NewShapeCommandApplier(eventsFramework pkg.Provider) internal.CommandApplier
}

func (f *shapeFacade) NewCommandHandlerBuilder() *cqrs.CommandHandlerBuilder[internal.CommandApplier] {
	return &cqrs.CommandHandlerBuilder[internal.CommandApplier]{}
}

func (f *shapeFacade) NewShapeCommandApplier(eventsFramework pkg.Provider) internal.CommandApplier {
	return internal.NewShapeCommandApplier(eventsFramework)
}
