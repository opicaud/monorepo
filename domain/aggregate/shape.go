package aggregate

type Shape interface {
	HandleNewShape(command newShapeCommand) ShapeCreated
	HandleStretchCommand(command newStretchCommand) ShapeStreched
	ApplyShapeCreatedEvent(area ShapeCreated) Shape
	ApplyShapeStrechedEvent(area ShapeStreched) Shape
}
