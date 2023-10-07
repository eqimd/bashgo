package splitter_test

import (
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/env"
	"github.com/eqimd/bashgo/internal/command/external"
	"github.com/eqimd/bashgo/internal/command/splitter"
	"github.com/stretchr/testify/require"
)

func TestParseEnvCommand(t *testing.T) {
	splitter := splitter.NewCommandSplitterImpl()

	envName := "abc"
	envValue := "def"
	line := envName + "=" + envValue

	com, args, err := splitter.Split(line)

	require.NoError(t, err, "unexpected error: %v")
	require.Equal(t,
		env.NewEnvVariableCommand(envName, envValue),
		com,
	)
	require.Equal(t,
		make([]command.CommandArgument, 0),
		args,
	)
}

func TestSimpleExternalCommand(t *testing.T) {
	splitter := splitter.NewCommandSplitterImpl()

	line := "git --version"

	com, args, err := splitter.Split(line)

	require.NoError(t, err, "unexpected error: %v")

	expectedArgs := make([]command.CommandArgument, 1)
	expectedArgs[0] = "--version"
	require.Equal(
		t,
		external.NewExternalCommand("git"),
		com,
	)
	require.Equal(
		t,
		expectedArgs,
		args,
	)
}

func TestExternalCommandWithTwoSimpleArgs(t *testing.T) {
	splitter := splitter.NewCommandSplitterImpl()
	arg1 := "adasadsads"
	arg2 := "afvsfdffad"

	line := "ext " + arg1 + " " + arg2

	com, args, err := splitter.Split(line)

	require.NoError(t, err)

	expectedArgs := make([]command.CommandArgument, 2)
	expectedArgs[0] = command.CommandArgument(arg1)
	expectedArgs[1] = command.CommandArgument(arg2)
	require.Equal(
		t,
		external.NewExternalCommand("ext"),
		com,
	)
	require.Equal(
		t,
		expectedArgs,
		args,
	)
}

func TestExternalCommandWithQuotedArg(t *testing.T) {
	splitter := splitter.NewCommandSplitterImpl()
	arg1 := "sddsadsad asdsdasd"

	line := "ext '" + arg1 + "'"

	com, args, err := splitter.Split(line)

	require.NoError(t, err)

	expectedArgs := make([]command.CommandArgument, 2)
	expectedArgs[0] = command.CommandArgument(arg1)
	require.Equal(
		t,
		external.NewExternalCommand("ext"),
		com,
	)
	require.Equal(
		t,
		expectedArgs,
		args,
	)
}

func TestExternalCommandWithDoubleQuotesInsideSingleQuotedArg(t *testing.T) {
	splitter := splitter.NewCommandSplitterImpl()
	arg1 := "\"sddsadsad asdsdasd\""

	line := "ext '" + arg1 + "'"

	com, args, err := splitter.Split(line)

	require.NoError(t, err)

	expectedArgs := make([]command.CommandArgument, 2)
	expectedArgs[0] = command.CommandArgument(arg1)
	require.Equal(
		t,
		external.NewExternalCommand("ext"),
		com,
	)
	require.Equal(
		t,
		expectedArgs,
		args,
	)
}

func TestExternalCommandWithDoubleQuotedArg(t *testing.T) {
	splitter := splitter.NewCommandSplitterImpl()
	arg1 := "sddsadsad asdsdasd"

	line := "ext \"" + arg1 + "\""

	com, args, err := splitter.Split(line)

	require.NoError(t, err)

	expectedArgs := make([]command.CommandArgument, 2)
	expectedArgs[0] = command.CommandArgument(arg1)
	require.Equal(
		t,
		external.NewExternalCommand("ext"),
		com,
	)
	require.Equal(
		t,
		expectedArgs,
		args,
	)
}

func TestExternalCommandWithSingleQuotesInsideDoubleQuotedArg(t *testing.T) {
	splitter := splitter.NewCommandSplitterImpl()
	arg1 := "sdds'adsad asds'dasd"

	line := "ext \"" + arg1 + "\""

	com, args, err := splitter.Split(line)

	require.NoError(t, err)

	expectedArgs := make([]command.CommandArgument, 2)
	expectedArgs[0] = command.CommandArgument(arg1)
	require.Equal(
		t,
		external.NewExternalCommand("ext"),
		com,
	)
	require.Equal(
		t,
		expectedArgs,
		args,
	)
}
