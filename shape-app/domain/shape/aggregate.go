package shape

import "github.com/opicaud/monorepo/events/pkg"

type Shape interface {
	HandleNewShape(command CreationCommand) Created
	HandleStretchCommand(command StretchCommand) Stretched
	ApplyShapeCreatedEvent(area Created) Shape
	ApplyShapeStretchedEvent(area Stretched) Shape
}

type Command[T CommandApplier] interface {
	Execute(apply T) ([]pkg.DomainEvent, error)
}

type CommandApplier interface {
	ApplyNewShapeCommand(command CreationCommand) ([]pkg.DomainEvent, error)
	ApplyNewStretchCommand(command StretchCommand) ([]pkg.DomainEvent, error)
}
