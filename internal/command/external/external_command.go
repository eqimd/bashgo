package external

import (
	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type ExternalCommand struct {
	executablePath string
}

func (cmd *ExternalCommand) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}

func NewExternalCommand(executablePath string) *ExternalCommand {
	return &ExternalCommand{executablePath}
}
