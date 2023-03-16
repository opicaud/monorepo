package shape

import "github.com/opicaud/monorepo/events/pkg"

type Shape interface {
	HandleNewShape(command newShapeCommand) Created
	HandleStretchCommand(command newStretchCommand) Stretched
	ApplyShapeCreatedEvent(area Created) Shape
	ApplyShapeStretchedEvent(area Stretched) Shape
}

type Command interface {
	Execute(apply ShapeCommandApplier) ([]pkg.DomainEvent, error)
}

type CommandHandler[T any] interface {
	Execute(command Command, applier T) error
}

type ShapeCommandApplier interface {
	ApplyNewShapeCommand(command newShapeCommand) ([]pkg.DomainEvent, error)
	ApplyNewStretchCommand(command newStretchCommand) ([]pkg.DomainEvent, error)
}
