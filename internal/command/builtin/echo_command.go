package builtin

import (
	"strings"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type echoRunner struct{}

func (r *echoRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	var strargs = make([]string, 0, len(args))
	for _, arg := range args {
		strargs = append(strargs, string(arg))
	}
	return &data.Output{Data: strings.Join(strargs, " "), ExitCode: 0}, nil
}
