package builtin

import (
	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type exitRunner struct{}

func (r *exitRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	return &data.Output{}, ErrExit
}
