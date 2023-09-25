package splitter

import "github.com/eqimd/bashgo/internal/command"

var _ CommandSplitter = &CommandSplitterImpl{}

type CommandSplitterImpl struct{}

func (splitter *CommandSplitterImpl) Split(s string) (command.Command, []command.CommandArgument, error) {
	// TODO
	return nil, nil, nil
}

func NewCommandSplitterImpl() *CommandSplitterImpl {
	return &CommandSplitterImpl{}
}
