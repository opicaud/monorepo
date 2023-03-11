package shape

import (
	"github.com/opicaud/monorepo/shape-app/eventstore"
)

type Shape interface {
	HandleNewShape(command newShapeCommand) Created
	HandleStretchCommand(command newStretchCommand) Stretched
	ApplyShapeCreatedEvent(area Created) Shape
	ApplyShapeStretchedEvent(area Stretched) Shape
}

type Command interface {
	Execute(apply ApplyShapeCommand) ([]eventstore.DomainEvent, error)
}

type CommandHandler interface {
	Execute(command Command) error
}

type ApplyShapeCommand interface {
	ApplyNewShapeCommand(command newShapeCommand) ([]eventstore.DomainEvent, error)
	ApplyNewStretchCommand(command newStretchCommand) ([]eventstore.DomainEvent, error)
}
