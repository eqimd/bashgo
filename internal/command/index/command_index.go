package index

import "github.com/eqimd/bashgo/internal/command"

type commandIndex struct{}

func (index *commandIndex) LookupCommand(cmd string) (command.Command, error) {
	// TODO
	return nil, nil
}

var CommandIndex = &commandIndex{}
