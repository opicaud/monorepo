package shape

import (
	"example2/domain/adapter"
)

type Shape interface {
	HandleNewShape(command newShapeCommand) Created
	HandleStretchCommand(command newStretchCommand) Stretched
	ApplyShapeCreatedEvent(area Created) Shape
	ApplyShapeStretchedEvent(area Stretched) Shape
}

type Command interface {
	Execute(apply ApplyShapeCommand) ([]adapter.DomainEvent, error)
}

type CommandHandler interface {
	Execute(command Command) error
}

type ApplyShapeCommand interface {
	ApplyNewShapeCommand(command newShapeCommand) ([]adapter.DomainEvent, error)
	ApplyNewStretchCommand(command newStretchCommand) ([]adapter.DomainEvent, error)
}
