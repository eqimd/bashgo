package command

import (
	"github.com/eqimd/bashgo/internal/data"
)

type CommandArgument string

/*
 * Интерфейс, описывающий команду, которую можно запустить
 */
type Command interface {
	/*
	 * Метод для запуска команды
	 */
	Run(args []CommandArgument, input *data.Input) (*data.Output, error)
}
