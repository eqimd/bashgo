package bash

import (
	"fmt"

	"github.com/eqimd/bashgo/internal/pipeline/parser"
)

var _ Bash = &BashImpl{}

/*
 * Реализация интерфейса Bash
 */
type BashImpl struct {
	pipelineParser parser.PipelineParser
}

func (bash *BashImpl) Execute(command string) (string, int, error) {
	pipeline := bash.pipelineParser.Parse(command)

	output, err := pipeline.Run(nil)
	if err != nil {
		return "", 0, fmt.Errorf("can't run pipeline: %w", err)
	}

	return output.Data, output.ExitCode, nil
}

/*
 * Конструктор для BashImpl. Принимает на вход PipeParser
 */
func NewBashImpl(pipelineParser parser.PipelineParser) *BashImpl {
	return &BashImpl{pipelineParser}
}
