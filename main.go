package main

import (
	"github.com/eqimd/bashgo/internal/bash"
	"github.com/eqimd/bashgo/internal/command/splitter"
	"github.com/eqimd/bashgo/internal/config"
	"github.com/eqimd/bashgo/internal/pipe/parser"
	"github.com/eqimd/bashgo/internal/repl"
	"github.com/spf13/pflag"
)

var isDebug bool

func main() {
	pflag.BoolVar(&isDebug, "debug", false, "enable debug information")

	config.InitConfig(isDebug)

	cmdSplitter := splitter.NewCommandSplitterImpl()
	pipeParser := parser.NewPipeParserImpl(cmdSplitter)
	bash := bash.NewBashImpl(pipeParser)
	repl := repl.NewRepl(bash)

	if err := repl.StartRepl(); err != nil {
		panic(err)
	}
}
