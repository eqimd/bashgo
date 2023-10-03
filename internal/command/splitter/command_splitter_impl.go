package splitter

import "github.com/eqimd/bashgo/internal/command"
import "strings"
import "github.com/eqimd/bashgo/internal/command/builtin"
import "github.com/eqimd/bashgo/internal/command/external"

var _ CommandSplitter = &CommandSplitterImpl{}

type CommandSplitterImpl struct{}

func (splitter *CommandSplitterImpl) Split(s string) (command.Command, []command.CommandArgument, error) {
	words := strings.Split(s, " ")
	var commandWord = words[0]

	args := make([]command.CommandArgument, len(words[1:]))
	for i, w := range words[1:] {
		args[i] = command.CommandArgument(w)
	}

	com, err := builtin.LookupBuiltinCommand(commandWord)

	if err == nil {
		return com, args, nil
	}
	return external.NewExternalCommand(commandWord), args, nil
	return nil, nil, nil
}

func NewCommandSplitterImpl() *CommandSplitterImpl {
	return &CommandSplitterImpl{}
}
