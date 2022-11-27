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

func (f *fullShapeCommandHandler) Handle(command Command) error {
	err := command.Execute()
	f.repository.Save(command.(fullShapeCommand).shape)
	return err
}
