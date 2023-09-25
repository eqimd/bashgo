package env

import (
	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type EnvVariableCommand struct {
	name  string
	value string
}

func (cmd *EnvVariableCommand) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}

func NewEnvVariableCommand(name string, value string) *EnvVariableCommand {
	return &EnvVariableCommand{name, value}
}
