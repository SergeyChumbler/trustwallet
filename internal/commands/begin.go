package commands

import (
	"trustwallet/internal/store"
)

type beginCommand struct {
	name  string
	store store.Store
}

func NewBeginCommand(s store.Store) Command {
	d := &beginCommand{
		name:  "begin",
		store: s,
	}

	return d
}

func (g beginCommand) GetName() string {
	return g.name
}

func (g beginCommand) Validate(args []string) error {
	if len(args) != 0 {
		return ErrBadCommandFormat
	}

	return nil
}

func (g beginCommand) Handle(_ []string) OperationResult {
	g.store.Begin()

	return OperationResult{"", nil}
}
