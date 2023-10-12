package parser

import (
	"github.com/eqimd/bashgo/internal/command/splitter"
	"github.com/eqimd/bashgo/internal/pipe"
)

var _ PipeParser = &PipeParserImpl{}

type PipeParserImpl struct {
	splitter splitter.CommandSplitter
}

// Метод создаёт исполяемый Pipe из строки, описывающей его
func (parser *PipeParserImpl) Parse(s string) *pipe.Pipe {
	command, args := parser.splitter.Split(s)

	return pipe.NewPipe(command, args)
}

// Создание нового парсера
func NewPipeParserImpl(splitter splitter.CommandSplitter) *PipeParserImpl {
	return &PipeParserImpl{splitter}
}
