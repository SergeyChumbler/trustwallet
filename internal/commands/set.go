package commands

import (
	"trustwallet/internal/store"
)

type setCommand struct {
	name  string
	store store.Store
}

func NewSetCommand(s store.Store) Command {
	g := &setCommand{
		name:  "set",
		store: s,
	}

	return g
}

func (g setCommand) GetName() string {
	return g.name
}

func (g setCommand) Validate(args []string) error {
	if len(args) != 2 {
		return ErrBadCommandFormat
	}

	return nil
}

func (g setCommand) Handle(args []string) OperationResult {
	g.store.Set(args[0], args[1])

	return OperationResult{"", nil}
}
