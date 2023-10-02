package builtin_test

import (
	"os"
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/builtin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPwdNaive(t *testing.T) {
	testdir, err := os.Getwd()
	if err != nil {
		panic("unexpected error during testing")
	}
	output, err := builtin.Pwd.Run([]command.CommandArgument{}, nil)
	require.NoError(t, err, "unexpected error: %w", err)
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	assert.Equal(t, testdir, output.Data, "wrong directory returned: want %w, got %w", testdir, output.Data)
}
