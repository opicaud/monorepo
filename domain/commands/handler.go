package commands

import (
	"example2/domain/repository"
)

func NewHandlerCommand(repository repository.Repository) CommandHandler {
	standardCommandHandler := new(standardCommandHandler)
	standardCommandHandler.repository = repository
	return standardCommandHandler
}

type standardCommandHandler struct {
	repository repository.Repository
}

func (standardCommandHandler) Handler(command Command) error {
	err := command.Execute()
	return err
}
