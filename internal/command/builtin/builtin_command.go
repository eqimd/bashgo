package builtin

import (
	"errors"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

var (
	Cat  = &builtinCommand{"cat", &catRunner{}}
	Echo = &builtinCommand{"echo", &echoRunner{}}
	Wc   = &builtinCommand{"wc", &wcRunner{}}
)

var ErrNoCommand = errors.New("command does not exist")

type builtinCommand struct {
	Name string

	command.Command
}

type catRunner struct{}

func (r *catRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}

type echoRunner struct{}

func (r *echoRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}

type wcRunner struct{}

func (r *wcRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}

func LookupBuiltinCommand(cmd string) (*builtinCommand, error) {
	switch cmd {
	case Cat.Name:
		return Cat, nil
	case Echo.Name:
		return Echo, nil
	case Wc.Name:
		return Wc, nil
	default:
		return nil, ErrNoCommand
	}
}
