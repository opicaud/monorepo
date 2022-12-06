package aggregate

import "example2/infra"

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
