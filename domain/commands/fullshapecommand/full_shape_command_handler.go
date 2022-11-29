package fullshapecommand

import (
	"example2/domain/commands"
	"example2/domain/repository"
)

func NewFullShapeCommandHandler(repository repository.Repository) commands.CommandHandler {
	fullShapeCommandHandler := new(fullShapeCommandHandler)
	fullShapeCommandHandler.repository = repository
	return fullShapeCommandHandler
}

type fullShapeCommandHandler struct {
	repository repository.Repository
}

func (f *fullShapeCommandHandler) Execute(command commands.Command) error {
	shape := command.(fullShapeCommand).shape
	err := shape.Area()
	f.repository.Save(shape)
	return err
}
