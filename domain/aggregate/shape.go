package aggregate

import "example2/infra"

type Shape interface {
	HandleNewShape(command newShapeCommand) ShapeCreated
	HandleStretchCommand(command newStretchCommand) ShapeStreched
	ApplyShapeCreatedEvent(area ShapeCreated) Shape
	ApplyShapeStrechedEvent(area ShapeStreched) Shape
}

type ShapeCommand interface {
	Apply(apply ApplyShapeCommand) []infra.Event
}

type ShapeCommandHandler interface {
	Execute(command ShapeCommand) error
}

type ApplyShapeCommand interface {
	ApplyNewShapeCommand(command newShapeCommand) []infra.Event
	ApplyNewStretchCommand(command newStretchCommand) []infra.Event
}
