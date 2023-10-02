package builtin

import (
	"fmt"

	"github.com/eqimd/bashgo/internal/command"
)

var (
	Cat  = &builtinCommand{"cat", &catRunner{}}
	Echo = &builtinCommand{"echo", &echoRunner{}}
	Wc   = &builtinCommand{"wc", &wcRunner{}}
)

var ErrNoCommand = fmt.Errorf("command does not exist")

type builtinCommand struct {
	Name string

	command.Command
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
