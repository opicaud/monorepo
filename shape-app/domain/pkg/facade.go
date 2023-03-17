package pkg

import (
	"github.com/google/uuid"
	"github.com/opicaud/monorepo/events/pkg"
	"github.com/opicaud/monorepo/shape-app/cqrs"
	internal2 "github.com/opicaud/monorepo/shape-app/domain/internal"
)

func (f *shapeFacade) NewCreationShapeCommand(nature string, dimensions ...float32) internal2.Command[internal2.CommandApplier] {
	command := internal2.NewCreationShapeCommand(nature, dimensions)
	return command
}

func (f *shapeFacade) NewStretchShapeCommand(id uuid.UUID, stretchBy float32) internal2.Command[internal2.CommandApplier] {
	command := internal2.NewStretchShapeCommand(id, stretchBy)
	return command
}

func New() ShapeFacade {
	return new(shapeFacade)
}

type shapeFacade struct{}

type ShapeFacade interface {
	NewCreationShapeCommand(nature string, dimensions ...float32) internal2.Command[internal2.CommandApplier]
	NewStretchShapeCommand(id uuid.UUID, stretchBy float32) internal2.Command[internal2.CommandApplier]
	NewCommandHandlerBuilder() *cqrs.CommandHandlerBuilder[internal2.CommandApplier]
	NewShapeCommandApplier(eventsFramework pkg.Provider) internal2.CommandApplier
}

func (f *shapeFacade) NewCommandHandlerBuilder() *cqrs.CommandHandlerBuilder[internal2.CommandApplier] {
	return &cqrs.CommandHandlerBuilder[internal2.CommandApplier]{}
}

func (f *shapeFacade) NewShapeCommandApplier(eventsFramework pkg.Provider) internal2.CommandApplier {
	return internal2.NewShapeCommandApplier(eventsFramework)
}
