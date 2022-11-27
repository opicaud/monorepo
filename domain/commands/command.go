package commands

type Command interface{}

type CommandHandler interface {
	Execute(command Command) error
}
