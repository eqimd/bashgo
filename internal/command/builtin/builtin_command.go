package builtin

import (
	"fmt"

	"github.com/eqimd/bashgo/internal/command"
)

var (
	Cat  = &builtinCommand{"cat", &catRunner{}}
	Echo = &builtinCommand{"echo", &echoRunner{}}
	Exit = &builtinCommand{"exit", &exitRunner{}}
	Pwd  = &builtinCommand{"pwd", &pwdRunner{}}
	Wc   = &builtinCommand{"wc", &wcRunner{}}
)

var ErrExit = fmt.Errorf("exit initiated")
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
	case Exit.Name:
		return Exit, nil
	case Pwd.Name:
		return Pwd, nil
	case Wc.Name:
		return Wc, nil
	default:
		return nil, ErrNoCommand
	}
}
