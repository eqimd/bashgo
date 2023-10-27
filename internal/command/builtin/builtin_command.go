package builtin

import (
	"fmt"

	"github.com/eqimd/bashgo/internal/command"
)

/*
 * Реализации builtin-комманд
 */
var (
	builtinCommands = map[string]command.Command{
		"cat":  &catRunner{},
		"echo": &echoRunner{},
		"exit": &exitRunner{},
		"pwd":  &pwdRunner{},
		"wc":   &wcRunner{},
		"grep": &grepRunner{},
	}
)

var ErrExit = fmt.Errorf("exit initiated")
var ErrNoCommand = fmt.Errorf("command does not exist")

/*
 * Функция для получения встроенной команды.
 * Если функции с нужным именем не существует, возвращает ошибку
 */
func LookupBuiltinCommand(cmd string) (command.Command, error) {
	if com, ok := builtinCommands[cmd]; ok {
		return com, nil
	}

	return nil, ErrNoCommand
}
