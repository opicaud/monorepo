package valueobject

import "example2/infra"

type Shape interface {
	HandleNewShape(command newShapeCommand) infra.Event
}
