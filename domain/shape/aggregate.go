package shape

import "example2/infra"

type Shape interface {
	HandleNewShape(command newShapeCommand) ShapeCreated
	HandleStretchCommand(command newStretchCommand) ShapeStreched
	ApplyShapeCreatedEvent(area ShapeCreated) Shape
	ApplyShapeStrechedEvent(area ShapeStreched) Shape
}

type Command interface {
	Apply(apply ApplyShapeCommand) ([]infra.Event, error)
}

type CommandHandler interface {
	Execute(command Command) error
}

type ApplyShapeCommand interface {
	ApplyNewShapeCommand(command newShapeCommand) ([]infra.Event, error)
	ApplyNewStretchCommand(command newStretchCommand) ([]infra.Event, error)
}
