package bash

import (
	"fmt"

	"github.com/eqimd/bashgo/internal/pipe/parser"
)

var _ Bash = &BashImpl{}

/*
 * Реализация интерфейса Bash
 */
type BashImpl struct {
	pipeParser parser.PipeParser
}

func (bash *BashImpl) Execute(command string) (string, int, error) {
	pipe, err := bash.pipeParser.Parse(command)
	if err != nil {
		return "", 0, fmt.Errorf("can't parse pipe: %w", err)
	}

	output, err := pipe.RunPipe(nil)
	if err != nil {
		return "", 0, fmt.Errorf("can't run pipe: %w", err)
	}

	return output.Data, output.ExitCode, nil
}

/*
 * Конструктор для BashImpl. Принимает на вход PipeParser
 */
func NewBashImpl(pipeParser parser.PipeParser) *BashImpl {
	return &BashImpl{pipeParser}
}
