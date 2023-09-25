package splitter

import "github.com/eqimd/bashgo/internal/command"

type CommandSplitter interface {
	Split(s string) (command.Command, []command.CommandArgument, error)
}
