package builtin_test

import (
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/builtin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCatEmpty(t *testing.T) {
	filename := "file_empty"
	output, err := builtin.Cat.Run([]command.CommandArgument{command.CommandArgument(filename)}, nil)

	require.NoError(t, err, "unexpected error: %v")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	assert.Equal(t, "", output.Data, "wrong amount of values returned: want %v, got %v", "", output.Data)
}

func TestCatSingle(t *testing.T) {
	filename := "file_1"
	output, err := builtin.Cat.Run([]command.CommandArgument{command.CommandArgument(filename)}, nil)
	exp := "This is the first part."

	require.NoError(t, err, "unexpected error: %v")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	assert.Equal(t, exp, output.Data, "wrong amount of values returned: want %v, got %v", exp, output.Data)
}

func TestCatMultiple(t *testing.T) {
	filename1 := "file_1"
	filename2 := "file_2"
	output, err := builtin.Cat.Run([]command.CommandArgument{
		command.CommandArgument(filename1),
		command.CommandArgument(filename2),
	}, nil)
	exp := "This is the first part.\n.trap dnoces eht si sihT"

	require.NoError(t, err, "unexpected error: %v")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	assert.Equal(t, exp, output.Data, "wrong amount of values returned: want %v, got %v", exp, output.Data)
}

func TestCatSame(t *testing.T) {
	filename := "file_bytes"
	output, err := builtin.Cat.Run([]command.CommandArgument{
		command.CommandArgument(filename),
		command.CommandArgument(filename),
	}, nil)
	exp := "\x00\x01\x02\x03\x04\x05\n\x00\x01\x02\x03\x04\x05"

	require.NoError(t, err, "unexpected error: %v")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	assert.Equal(t, exp, output.Data, "wrong amount of values returned: want %v, got %v", exp, output.Data)
}
