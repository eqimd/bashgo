package main

import (
	"github.com/eqimd/bashgo/internal/bash"
	cmdsplitter "github.com/eqimd/bashgo/internal/command/splitter"
	"github.com/eqimd/bashgo/internal/config"
	pipeparser "github.com/eqimd/bashgo/internal/pipe/parser"
	pipelineparser "github.com/eqimd/bashgo/internal/pipeline/parser"
	pipelinesplitter "github.com/eqimd/bashgo/internal/pipeline/splitter"
	"github.com/eqimd/bashgo/internal/repl"
	"github.com/spf13/pflag"
)

var isDebug bool

func main() {
	pflag.BoolVar(&isDebug, "debug", false, "enable debug information")

	config.InitConfig(isDebug)

	cmdSplitter := cmdsplitter.NewCommandSplitterImpl()
	pipeParser := pipeparser.NewPipeParserImpl(cmdSplitter)

	pipelineSplitter := pipelinesplitter.NewPipelineSplitterImpl()
	pipelineParser := pipelineparser.NewPipelineParserImpl(pipelineSplitter, pipeParser)

	bash := bash.NewBashImpl(pipelineParser)
	repl := repl.NewRepl(bash)

	if err := repl.StartRepl(); err != nil {
		panic(err)
	}
}
