package command

import (
	"github.com/eqimd/bashgo/internal/data"
)

type CommandArgument string

type Command interface {
	Run(args []CommandArgument, input *data.Input) (*data.Output, error)
}
