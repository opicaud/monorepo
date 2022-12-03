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
	shape := loadShape(command.(newShapeCommand))
	applyCommandOnAggregate(command, shape)
	return f.repository.Save(shape)
}

func loadShape(command newShapeCommand) Shape {
	shape, err := newShapeBuilder().CreateAShape(command.nature).WithDimensions(command.dimensions)
	if err != nil {
		panic(err)
	}
	return shape
}

func applyCommandOnAggregate(command commands.Command, shape Shape) {
	switch v := command.(type) {
	default:
		fmt.Printf("unexpected command %T", v)
	case newShapeCommand:
		shape.HandleNewShape(command.(newShapeCommand))
	}
}
