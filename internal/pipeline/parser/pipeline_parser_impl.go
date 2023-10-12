package parser

import (
	"github.com/eqimd/bashgo/internal/pipe"
	pipeparser "github.com/eqimd/bashgo/internal/pipe/parser"
	"github.com/eqimd/bashgo/internal/pipeline"
	"github.com/eqimd/bashgo/internal/pipeline/splitter"
)

var _ PipelineParser = &PipelineParserImpl{}

/*
 * Структура для парсера пайпланов. Для работы ей нужен сплиттер пайпланов и парсер пайпов
 */
type PipelineParserImpl struct {
	splitter   splitter.PipelineSplitter
	pipeParser pipeparser.PipeParser
}

func (parser *PipelineParserImpl) Parse(cmd string) pipeline.Pipeline {
	cmdStrings := parser.splitter.Split(cmd)
	pipes := []*pipe.Pipe{}
	for _, cmd := range cmdStrings {
		pipes = append(pipes, parser.pipeParser.Parse(cmd))
	}

	pipeline := pipeline.NewPipelineImpl(pipes)
	return pipeline
}

func NewPipelineParserImpl(
	splitter splitter.PipelineSplitter,
	pipeParser pipeparser.PipeParser,
) *PipelineParserImpl {
	return &PipelineParserImpl{splitter: splitter}
}
