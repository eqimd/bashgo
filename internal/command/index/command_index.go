package index

import (
	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/builtin"
	"github.com/eqimd/bashgo/internal/command/external"
)

type commandIndex struct{}

/*
 * Ищет соответствие имени команды её класса и преобразует в соответствующий класс
 */
func (index *commandIndex) LookupCommand(cmd string) (command.Command, error) {
	com, err := builtin.LookupBuiltinCommand(cmd)

	if err == nil {
		return com, nil
	}

	return external.NewExternalCommand(cmd), nil
}

var CommandIndex = &commandIndex{}
