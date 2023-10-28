package builtin_test

import (
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/builtin"
	"github.com/eqimd/bashgo/internal/data"
	"github.com/stretchr/testify/require"
)

var Grep, _ = builtin.LookupBuiltinCommand("grep")

func TestGrep_Simple(t *testing.T) {
	input := &data.Input{Data: "Test text"}
	output, err := Grep.Run([]command.CommandArgument{command.CommandArgument("text")}, input)

	require.NoError(t, err, "no error expected")
	require.Equal(t, "Test text grep\n", output.Data)
}

func TestGrep_CaseInsensitive(t *testing.T) {
	input := &data.Input{Data: "Test text"}
	output, err := Grep.Run(
		[]command.CommandArgument{
			command.CommandArgument("-i"),
			command.CommandArgument("test"),
		},
		input,
	)

	require.NoError(t, err, "no error expected")
	require.Equal(t, "Test text grep\n", output.Data)
}

func TestGrep_WordPositive(t *testing.T) {
	input := &data.Input{Data: "Test text"}
	output, err := Grep.Run(
		[]command.CommandArgument{
			command.CommandArgument("-w"),
			command.CommandArgument("text"),
		},
		input,
	)

	require.NoError(t, err, "no error expected")
	require.Equal(t, "Test text grep\n", output.Data)
}

func TestGrep_WordNegative(t *testing.T) {
	input := &data.Input{Data: "Test text"}
	output, err := Grep.Run(
		[]command.CommandArgument{
			command.CommandArgument("-w"),
			command.CommandArgument("Tes"),
		},
		input,
	)

	require.NoError(t, err, "no error expected")
	require.Equal(t, "", output.Data)
}
