package builtin_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/builtin"
	"github.com/eqimd/bashgo/internal/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var basicInput = &data.Input{Data: "Hello, World!\nThis is bashgo!"}
var advancedInput = &data.Input{Data: "Hello, World 👋\nWe got emojis now 😎\n 🤟🙌👨🍊 🧊🧩\nИ русские слова тоже\n\n\n"}

func parseAndAssertOutput(
	t *testing.T,
	output *data.Output,
	needFilename bool,
	needLines bool,
	needWords bool,
	needChars bool,
	filenameRequired string,
	linesRequired int,
	wordsRequired int,
	charsRequired int,
) {
	res := strings.Split(output.Data, " ")
	var valuesRequired int
	if needFilename {
		valuesRequired++
	}
	if needLines {
		valuesRequired++
	}
	if needWords {
		valuesRequired++
	}
	if needChars {
		valuesRequired++
	}
	require.Equal(t, valuesRequired, len(res), "wrong amount of values returned: want %v, got %v", valuesRequired, len(res))

	var filename string
	var lines int
	var words int
	var chars int
	var curIdx int
	var err error
	if needFilename {
		filename = res[curIdx]
		curIdx++
	}
	if needLines {
		lines, err = strconv.Atoi(res[curIdx])
		require.NoError(t, err, "could not parse lines amount, got: %v", res[curIdx])
		curIdx++
	}
	if needWords {
		words, err = strconv.Atoi(res[curIdx])
		require.NoError(t, err, "could not parse words amount, got: %v", res[curIdx])
		curIdx++
	}
	if needChars {
		chars, err = strconv.Atoi(res[curIdx])
		require.NoError(t, err, "could not parse chars amount, got: %v", res[curIdx])
		curIdx++
	}

	if needFilename {
		assert.Equal(t, filenameRequired, filename, "wrong amount of lines found: want %v, got %v", filenameRequired, filename)
	}
	if needLines {
		assert.Equal(t, linesRequired, lines, "wrong amount of lines found: want %v, got %v", linesRequired, lines)
	}
	if needWords {
		assert.Equal(t, wordsRequired, words, "wrong amount of lines found: want %v, got %v", wordsRequired, words)
	}
	if needChars {
		assert.Equal(t, charsRequired, chars, "wrong amount of lines found: want %v, got %v", charsRequired, chars)
	}
}

func TestEmptyInput(t *testing.T) {
	output, err := builtin.Wc.Run([]command.CommandArgument{}, &data.Input{Data: ""})
	require.NoError(t, err, "unexpected error")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	parseAndAssertOutput(t, output, false, true, true, true, "", 0, 0, 0)
}

func TestBasicInput(t *testing.T) {
	output, err := builtin.Wc.Run([]command.CommandArgument{}, basicInput)
	require.NoError(t, err, "unexpected error")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	parseAndAssertOutput(t, output, false, true, true, true, "", 2, 5, 29)
}

func TestBasicInputWithFlags(t *testing.T) {
	output, err := builtin.Wc.Run([]command.CommandArgument{command.CommandArgument("-lw")}, basicInput)
	require.NoError(t, err, "unexpected error")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	parseAndAssertOutput(t, output, false, true, true, false, "", 2, 5, 0)
}

func TestInput(t *testing.T) {
	output, err := builtin.Wc.Run([]command.CommandArgument{}, advancedInput)
	require.NoError(t, err, "unexpected error")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	parseAndAssertOutput(t, output, false, true, true, true, "", 7, 14, 108)
}

func TestEmptyFile(t *testing.T) {
	var filename = "file_empty"
	output, err := builtin.Wc.Run([]command.CommandArgument{command.CommandArgument(filename)}, nil)
	require.NoError(t, err, "unexpected error")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	parseAndAssertOutput(t, output, true, true, true, true, filename, 0, 0, 0)
}

func TestFile(t *testing.T) {
	var filename = "file_advanced"
	output, err := builtin.Wc.Run([]command.CommandArgument{command.CommandArgument(filename)}, nil)
	require.NoError(t, err, "unexpected error")
	require.Equal(t, 0, output.ExitCode, "non-zero exitcode")
	parseAndAssertOutput(t, output, true, true, true, true, filename, 7, 14, 108)
}
