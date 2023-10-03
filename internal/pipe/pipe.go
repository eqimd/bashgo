package pipe

import (
	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type Pipe struct {
	command   command.Command
	arguments []command.CommandArgument
}

func (pipe *Pipe) RunPipe(input *data.Input) (*data.Output, error) {
	return nil, nil
}

func NewPipe(command command.Command, arguments []command.CommandArgument) *Pipe {
	return &Pipe{command, arguments}
}
