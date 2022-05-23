package commands

import (
	"trustwallet/internal/store"
)

type getCommand struct {
	name  string
	store store.Store
}

func NewGetCommand(s store.Store) Command {
	g := &getCommand{
		name:  "get",
		store: s,
	}

	return g
}

func (g getCommand) GetName() string {
	return g.name
}

func (g getCommand) Validate(args []string) error {
	if len(args) != 1 {
		return ErrBadCommandFormat
	}

	return nil
}

func (g getCommand) Handle(args []string) OperationResult {
	val, err := g.store.Get(args[0])

	return OperationResult{val, err}
}
