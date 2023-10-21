package builtin

import (
	"os"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type pwdRunner struct{}

func (r *pwdRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	dir, err := os.Getwd()
	if err != nil {
		return &data.Output{ExitCode: 255}, err
	}
	return &data.Output{Data: dir}, nil
}
