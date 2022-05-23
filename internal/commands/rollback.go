package commands

import (
	"trustwallet/internal/store"
)

type rollbackCommand struct {
	name  string
	store store.Store
}

func NewRollbackCommand(s store.Store) Command {
	d := &rollbackCommand{
		name:  "rollback",
		store: s,
	}

	return d
}

func (g rollbackCommand) GetName() string {
	return g.name
}

func (g rollbackCommand) Validate(args []string) error {
	if len(args) != 0 {
		return ErrBadCommandFormat
	}

	return nil
}

func (g rollbackCommand) Handle(_ []string) OperationResult {
	err := g.store.Rollback()

	return OperationResult{"", err}
}
