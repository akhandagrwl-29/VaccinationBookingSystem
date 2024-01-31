package model

type CMD interface {
	Execute(command *Command) error
}

type Command struct {
	CommandName string
	Connection  CMD
	Arguments   interface{}
}
