package aggregate

type Shape interface {
	HandleCaculateShapeArea(command newShapeCommand) AreaShapeCalculated
}
