package commands

import (
	"trustwallet/internal/store"
)

type commitCommand struct {
	name  string
	store store.Store
}

func NewCommitCommand(s store.Store) Command {
	d := &commitCommand{
		name:  "commit",
		store: s,
	}

	return d
}

func (g commitCommand) GetName() string {
	return g.name
}

func (g commitCommand) Validate(args []string) error {
	if len(args) != 0 {
		return ErrBadCommandFormat
	}

	return nil
}

func (g commitCommand) Handle(_ []string) OperationResult {
	err := g.store.Commit()

	return OperationResult{"", err}
}
