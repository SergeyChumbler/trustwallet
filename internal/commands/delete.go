package commands

import (
	"trustwallet/internal/store"
)

type deleteCommand struct {
	name  string
	store store.Store
}

func NewDeleteCommand(s store.Store) Command {
	d := &deleteCommand{
		name:  "delete",
		store: s,
	}

	return d
}

func (g deleteCommand) GetName() string {
	return g.name
}

func (g deleteCommand) Validate(args []string) error {
	if len(args) != 1 {
		return ErrBadCommandFormat
	}

	return nil
}

func (g deleteCommand) Handle(args []string) OperationResult {
	err := g.store.Delete(args[0])

	return OperationResult{"", err}
}
