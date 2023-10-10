package splitter

import "github.com/eqimd/bashgo/internal/command"

type CommandSplitter interface {
	// Данный метод разделяет строку на команду и её аругменты, также приводя название команды к классу, ей соответствующему
	Split(s string) (command.Command, []command.CommandArgument, error)
}
