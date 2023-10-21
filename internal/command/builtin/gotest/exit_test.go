package builtin_test

import (
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/builtin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExitNaive(t *testing.T) {
	output, err := builtin.Exit.Run([]command.CommandArgument{}, nil)
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	assert.ErrorIs(t, err, builtin.ErrExit, "wrong error returned: want %v, got %v", err, builtin.ErrExit)
}
