package aggregate

import "example2/infra"

type ShapeCommand interface {
	Apply(apply ApplyShapeCommand) (Shape, []infra.Event)
}

type ShapeCommandHandler interface {
	Execute(command ShapeCommand) error
}

type ApplyShapeCommand interface {
	ApplyNewShapeCommand(command newShapeCommand) (Shape, []infra.Event)
	ApplyNewStretchCommand(command newStretchCommand) (Shape, []infra.Event)
}
