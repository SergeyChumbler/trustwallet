package commands

import (
	"strconv"
	"trustwallet/internal/store"
)

type countCommand struct {
	name  string
	store store.Store
}

func NewCountCommand(s store.Store) Command {
	d := &countCommand{
		name:  "count",
		store: s,
	}

	return d
}

func (g countCommand) GetName() string {
	return g.name
}

func (g countCommand) Validate(args []string) error {
	if len(args) != 1 {
		return ErrBadCommandFormat
	}

	return nil
}

func (g countCommand) Handle(args []string) OperationResult {
	count := g.store.Count(args[0])

	return OperationResult{strconv.Itoa(count), nil}
}
