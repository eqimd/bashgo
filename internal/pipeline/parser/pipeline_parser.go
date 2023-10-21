package parser

import "github.com/eqimd/bashgo/internal/pipeline"

type PipelineParser interface {
	Parse(cmd string) pipeline.Pipeline
}
