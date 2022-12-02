package shapecreationcommand

import (
	"example2/domain/commands"
	"example2/domain/repository"
)

func NewShapeCreationCommandHandler(repository repository.Repository) commands.CommandHandler {
	shapeCommandHandler := new(shapeCommandHandler)
	shapeCommandHandler.repository = repository
	return shapeCommandHandler
}

type shapeCommandHandler struct {
	repository repository.Repository
}

func (f *shapeCommandHandler) Execute(command commands.Command) error {
	shape := command.(newShapeCommand).shape
	shape.Execute(command)
	err := f.repository.Save(shape)
	return err
}
