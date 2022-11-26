package commands

type Command interface {
	Execute() error
}

type CommandHandler interface {
	Handler(command Command) error
}
