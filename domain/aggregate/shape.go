package aggregate

type Shape interface {
	HandleNewShape(command newShapeCommand) AreaShapeCalculated
	HandleStretchCommand(command newStretchCommand) AreaShapeCalculated
	ApplyShapeCreatedEvent(area ShapeCreatedEvent) Shape
	ApplyAreaShapeCalculated(area AreaShapeCalculated) Shape
}
