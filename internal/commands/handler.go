package commands

import (
	"errors"
	"trustwallet/internal/store"
)

var ErrUnsupportedCommand = errors.New("unsupported command")
var ErrBadCommandFormat = errors.New("bad command format")

type Command interface {
	GetName() string
	Validate(args []string) error
	Handle(args []string) OperationResult
}

type OperationResult struct {
	Data  string
	Error error
}

type CommandHandler interface {
	Handle(command string, args []string) OperationResult
}

type commandHandler struct {
	commands map[string]Command
}

func NewCommandHandler(s store.Store) CommandHandler {
	supportedCommands := []Command{
		NewGetCommand(s),
		NewSetCommand(s),
		NewDeleteCommand(s),
		NewCountCommand(s),
		NewBeginCommand(s),
		NewCommitCommand(s),
		NewRollbackCommand(s),
	}

	c := commandHandler{map[string]Command{}}
	for _, command := range supportedCommands {
		cmd := command
		c.commands[command.GetName()] = cmd
	}

	return c
}

func (ch commandHandler) Handle(command string, args []string) OperationResult {
	if c, ok := ch.commands[command]; ok {
		err := c.Validate(args)
		if err != nil {
			return OperationResult{
				Data:  "",
				Error: ErrBadCommandFormat,
			}
		}

		return c.Handle(args)
	}

	return OperationResult{
		Data:  "",
		Error: ErrUnsupportedCommand,
	}
}
