package gotest

import (
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/external"
	"github.com/eqimd/bashgo/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestExternalCommand(t *testing.T) {
	cmd := external.NewExternalCommand("echo")
	outp, err := cmd.Run([]command.CommandArgument{command.CommandArgument("testecho")}, nil)

	assert.NoError(t, err, "no error expected")

	assert.Equal(t, "testecho\n", outp.Data)

	assert.Equal(t, 0, outp.ExitCode)
}

func TestExternalCommand_NotInPath(t *testing.T) {
	cmd := external.NewExternalCommand("abra")
	_, err := cmd.Run([]command.CommandArgument{command.CommandArgument("1")}, nil)

	assert.Error(t, err, "expected not found in PATH")
}

func TestExternalCommand_Input(t *testing.T) {
	cmd := external.NewExternalCommand("cat")
	outp, err := cmd.Run([]command.CommandArgument{}, &data.Input{Data: "testcat"})

	assert.NoError(t, err, "no error expected")

	assert.Equal(t, "testcat", outp.Data)

	assert.Equal(t, 0, outp.ExitCode)
}
