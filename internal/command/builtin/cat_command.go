package builtin

import (
	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type catRunner struct{}

func (r *catRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}
