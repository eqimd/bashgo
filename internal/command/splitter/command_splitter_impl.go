package splitter

import (
	"strings"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/env"
	"github.com/eqimd/bashgo/internal/command/index"
)

var _ CommandSplitter = &CommandSplitterImpl{}

type CommandSplitterImpl struct{}

// Данный метод разделяет строку на команду и её аругменты, также приводя название команды к классу, ей соответствующему
func (splitter *CommandSplitterImpl) Split(s string) (command.Command, []command.CommandArgument, error) {
	var waitFor = rune(' ')
	var startFrom = 0
	words := make([]string, 0)
	s = s + " "
	for i, ch := range s {
		if waitFor == rune(' ') {
			if ch == rune('"') || ch == rune('\'') {
				startFrom = i + 1
				waitFor = ch
				continue
			}
		}
		if ch == waitFor {
			words = append(words, s[startFrom:i])
			startFrom = i + 1
			waitFor = rune(' ')
		}
	}
	var commandWord = words[0]

	args := make([]command.CommandArgument, len(words[1:]))
	for i, w := range words[1:] {
		args[i] = command.CommandArgument(w)
	}

	tryEnvSplit := strings.Split(commandWord, "=")
	if len(tryEnvSplit) == 2 {
		return env.NewEnvVariableCommand(tryEnvSplit[0], tryEnvSplit[1]), args, nil
	}

	com, err := index.CommandIndex.LookupCommand(commandWord)

	if err == nil {
		return com, args, nil
	} else {
		return nil, nil, err
	}
}

func NewCommandSplitterImpl() *CommandSplitterImpl {
	return &CommandSplitterImpl{}
}
