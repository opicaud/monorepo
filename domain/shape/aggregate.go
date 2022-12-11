package shape

import "example2/infra"

type Shape interface {
	HandleNewShape(command newShapeCommand) Created
	HandleStretchCommand(command newStretchCommand) Stretched
	ApplyShapeCreatedEvent(area Created) Shape
	ApplyShapeStretchedEvent(area Stretched) Shape
}

type Command interface {
	Execute(apply ApplyShapeCommand) ([]infra.Event, error)
}

type CommandHandler interface {
	Execute(command Command) error
}

type ApplyShapeCommand interface {
	ApplyNewShapeCommand(command newShapeCommand) ([]infra.Event, error)
	ApplyNewStretchCommand(command newStretchCommand) ([]infra.Event, error)
}
