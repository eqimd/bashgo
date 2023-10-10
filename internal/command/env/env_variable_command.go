package env

import (
	"os"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

/*
 * Структура, описывающая команду для задания переменной окружения, вида name=value
 */
type EnvVariableCommand struct {
	name  string
	value string
}

func (cmd *EnvVariableCommand) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	if err := os.Setenv(cmd.name, cmd.value); err != nil {
		return nil, err
	}
	return &data.Output{}, nil
}

/*
 * Конструктор для структуры переменной окружения
 */
func NewEnvVariableCommand(name string, value string) *EnvVariableCommand {
	return &EnvVariableCommand{name, value}
}
