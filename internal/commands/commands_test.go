package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trustwallet/internal/store"
)

func Test_GetCommandValidation(t *testing.T) {
	command := NewGetCommand(store.NewStore())

	type args struct {
		arguments []string
		valid     bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid zero args",
			args: args{
				arguments: []string{},
				valid:     false,
			},
		},
		{
			name: "invalid two args",
			args: args{
				arguments: []string{"a", "b"},
				valid:     false,
			},
		}, {
			name: "valid args",
			args: args{
				arguments: []string{"x"},
				valid:     true,
			},
		},
	}

	name := command.GetName()
	assert.Equal(t, name, "get")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := command.Validate(tt.args.arguments)
			if tt.args.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_SetCommandValidation(t *testing.T) {
	command := NewSetCommand(store.NewStore())

	type args struct {
		arguments []string
		valid     bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid zero args",
			args: args{
				arguments: []string{},
				valid:     false,
			},
		},
		{
			name: "invalid single args",
			args: args{
				arguments: []string{"a"},
				valid:     false,
			},
		},
		{
			name: "invalid three args",
			args: args{
				arguments: []string{"a", "b", "c"},
				valid:     false,
			},
		}, {
			name: "valid args",
			args: args{
				arguments: []string{"x", "y"},
				valid:     true,
			},
		},
	}

	name := command.GetName()
	assert.Equal(t, name, "set")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := command.Validate(tt.args.arguments)
			if tt.args.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_BeginCommandValidation(t *testing.T) {
	command := NewBeginCommand(store.NewStore())

	type args struct {
		arguments []string
		valid     bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid no zero args",
			args: args{
				arguments: []string{"a", "b"},
				valid:     false,
			},
		}, {
			name: "valid args",
			args: args{
				arguments: []string{},
				valid:     true,
			},
		},
	}

	name := command.GetName()
	assert.Equal(t, name, "begin")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := command.Validate(tt.args.arguments)
			if tt.args.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_CommitCommandValidation(t *testing.T) {
	command := NewCommitCommand(store.NewStore())

	type args struct {
		arguments []string
		valid     bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid none zero args",
			args: args{
				arguments: []string{"a"},
				valid:     false,
			},
		}, {
			name: "valid args",
			args: args{
				arguments: []string{},
				valid:     true,
			},
		},
	}

	name := command.GetName()
	assert.Equal(t, name, "commit")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := command.Validate(tt.args.arguments)
			if tt.args.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_CountCommandValidation(t *testing.T) {
	command := NewCountCommand(store.NewStore())

	type args struct {
		arguments []string
		valid     bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid zero args",
			args: args{
				arguments: []string{},
				valid:     false,
			},
		},
		{
			name: "invalid two args",
			args: args{
				arguments: []string{"a", "b"},
				valid:     false,
			},
		}, {
			name: "valid args",
			args: args{
				arguments: []string{"x"},
				valid:     true,
			},
		},
	}

	name := command.GetName()
	assert.Equal(t, name, "count")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := command.Validate(tt.args.arguments)
			if tt.args.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_DeleteCommandValidation(t *testing.T) {
	command := NewDeleteCommand(store.NewStore())

	type args struct {
		arguments []string
		valid     bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid zero args",
			args: args{
				arguments: []string{},
				valid:     false,
			},
		},
		{
			name: "invalid two args",
			args: args{
				arguments: []string{"a", "b"},
				valid:     false,
			},
		}, {
			name: "valid args",
			args: args{
				arguments: []string{"x"},
				valid:     true,
			},
		},
	}

	name := command.GetName()
	assert.Equal(t, name, "delete")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := command.Validate(tt.args.arguments)
			if tt.args.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_RollbackCommandValidation(t *testing.T) {
	command := NewRollbackCommand(store.NewStore())

	type args struct {
		arguments []string
		valid     bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid none zero args",
			args: args{
				arguments: []string{"a"},
				valid:     false,
			},
		}, {
			name: "valid args",
			args: args{
				arguments: []string{},
				valid:     true,
			},
		},
	}

	name := command.GetName()
	assert.Equal(t, name, "rollback")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := command.Validate(tt.args.arguments)
			if tt.args.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
