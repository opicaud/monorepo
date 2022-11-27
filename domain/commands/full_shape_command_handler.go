package commands

import (
	"example2/domain/repository"
)

func NewFullShapeCommandHandler(repository repository.Repository) CommandHandler {
	fullShapeCommandHandler := new(fullShapeCommandHandler)
	fullShapeCommandHandler.repository = repository
	return fullShapeCommandHandler
}

type fullShapeCommandHandler struct {
	repository repository.Repository
}

func (f *fullShapeCommandHandler) Execute(command Command) error {
	shape := command.(fullShapeCommand).shape
	err := shape.Area()
	f.repository.Save(shape)
	return err
}
