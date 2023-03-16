package shape

import "github.com/opicaud/monorepo/shape-app/events/pkg"

type Shape interface {
	HandleNewShape(command newShapeCommand) Created
	HandleStretchCommand(command newStretchCommand) Stretched
	ApplyShapeCreatedEvent(area Created) Shape
	ApplyShapeStretchedEvent(area Stretched) Shape
}

type Command interface {
	Execute(apply ApplyShapeCommand) ([]pkg.DomainEvent, error)
}

type CommandHandler interface {
	Execute(command Command) error
}

type ApplyShapeCommand interface {
	ApplyNewShapeCommand(command newShapeCommand) ([]pkg.DomainEvent, error)
	ApplyNewStretchCommand(command newStretchCommand) ([]pkg.DomainEvent, error)
}
