package valueobject

import "example2/domain/commands"

type Shape interface {
	CalculateArea()
	GetArea() float32
	commands.CommandHandler
}
