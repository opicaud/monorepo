package valueobject

import "example2/domain/commands"

type Shape interface {
	Area()
	GetArea() float32
	commands.CommandHandler
}
