package pipe

import (
	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

// Тип, описывающий одиночный Pipe -- команду и её аргументы
type Pipe struct {
	command   command.Command
	arguments []command.CommandArgument
}

// Метод запускает одиночный пайплайн, возвращая ответ и код ответа или ошибку
func (pipe *Pipe) RunPipe(input *data.Input) (*data.Output, error) {
	return pipe.command.Run(pipe.arguments, input)
}

// Метод для создания Pipe
func NewPipe(command command.Command, arguments []command.CommandArgument) *Pipe {
	return &Pipe{command, arguments}
}
