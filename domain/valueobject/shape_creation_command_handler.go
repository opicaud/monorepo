package valueobject

import (
	"example2/domain/commands"
	"fmt"
)

func NewShapeCreationCommandHandler(repository Repository) commands.CommandHandler {
	shapeCommandHandler := new(shapeCommandHandler)
	shapeCommandHandler.repository = repository
	return shapeCommandHandler
}

type shapeCommandHandler struct {
	repository Repository
}

func (f *shapeCommandHandler) Execute(command commands.Command) error {
	shape := command.(newShapeCommand).shape
	applyCommandOnAggregate(command, shape)
	err := f.repository.Save(shape)
	return err
}

func applyCommandOnAggregate(command commands.Command, shape Shape) {
	switch v := command.(type) {
	default:
		fmt.Printf("unexpected command %T", v)
	case newShapeCommand:
		shape.HandleNewShape(command.(newShapeCommand))
	}
}
