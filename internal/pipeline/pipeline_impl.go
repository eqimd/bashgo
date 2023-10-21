package pipeline

import (
	"errors"

	"github.com/eqimd/bashgo/internal/data"
	"github.com/eqimd/bashgo/internal/pipe"
)

var _ Pipeline = &PipelineImpl{}

/*
 * Структура для пайплайна. Хранит в себе пайпы для последовательного запуска
 */
type PipelineImpl struct {
	pipes []*pipe.Pipe
}

func (pipeline *PipelineImpl) Run(input *data.Input) (*data.Output, error) {
	curInp := input
	for _, pipe := range pipeline.pipes {
		outp, err := pipe.RunPipe(curInp)
		if err != nil {
			return nil, errors.Join(errors.New("pipeline run failed"), err)
		}

		if outp.ExitCode != 0 {
			return outp, nil
		}

		if curInp == nil {
			curInp = &data.Input{}
		}

		curInp.Data = outp.Data
	}

	outp := &data.Output{
		Data:     curInp.Data,
		ExitCode: 0,
	}

	return outp, nil
}

func NewPipelineImpl(pipes []*pipe.Pipe) *PipelineImpl {
	return &PipelineImpl{pipes: pipes}
}
