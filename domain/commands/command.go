package commands

type Command interface {
	Execute() error
}

type CommandHandler interface {
	Handle(command Command) error
}
