package builtin

import (
	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type echoRunner struct{}

func (r *echoRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}
