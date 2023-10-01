package builtin_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/builtin"
	"github.com/eqimd/bashgo/internal/data"
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
	if len(res) != valuesRequired {
		t.Fatalf("wrong amount of values returned: want %v, got %v", valuesRequired, len(res))
	}

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
		if err != nil {
			t.Fatalf("could not parse lines amount, got: %v", res[curIdx])
		}
		curIdx++
	}
	if needWords {
		words, err = strconv.Atoi(res[curIdx])
		if err != nil {
			t.Fatalf("could not parse words amount, got: %v", res[curIdx])
		}
		curIdx++
	}
	if needChars {
		chars, err = strconv.Atoi(res[curIdx])
		if err != nil {
			t.Fatalf("could not parse chars amount, got: %v", res[curIdx])
		}
		curIdx++
	}

	if needFilename && filename != filenameRequired {
		t.Errorf("wrong amount of lines found: want %v, got %v", linesRequired, lines)
	}
	if needLines && lines != linesRequired {
		t.Errorf("wrong amount of lines found: want %v, got %v", linesRequired, lines)
	}
	if needWords && words != wordsRequired {
		t.Errorf("wrong amount of words found: want %v, got %v", wordsRequired, words)
	}
	if needChars && chars != charsRequired {
		t.Errorf("wrong amount of chars found: want %v, got %v", charsRequired, chars)
	}
}

func TestWcEmptyInput(t *testing.T) {
	output, err := builtin.Wc.Run([]command.CommandArgument{}, &data.Input{Data: ""})
	if err != nil || output.ExitCode != 0 {
		t.Fatalf("non-zero exitcode")
	}
	parseAndAssertOutput(t, output, false, true, true, true, "", 0, 0, 0)
}

func TestWcBasicInput(t *testing.T) {
	output, err := builtin.Wc.Run([]command.CommandArgument{}, basicInput)
	if err != nil || output.ExitCode != 0 {
		t.Fatalf("non-zero exitcode")
	}
	parseAndAssertOutput(t, output, false, true, true, true, "", 2, 5, 29)
}

func TestWcBasicInputWithFlags(t *testing.T) {
	output, err := builtin.Wc.Run([]command.CommandArgument{command.CommandArgument("-lw")}, basicInput)
	if err != nil || output.ExitCode != 0 {
		t.Fatalf("non-zero exitcode")
	}
	parseAndAssertOutput(t, output, false, true, true, false, "", 2, 5, 0)
}

func TestWcInput(t *testing.T) {
	output, err := builtin.Wc.Run([]command.CommandArgument{}, advancedInput)
	if err != nil || output.ExitCode != 0 {
		t.Fatalf("non-zero exitcode")
	}
	parseAndAssertOutput(t, output, false, true, true, true, "", 7, 14, 108)
}

func TestWcEmptyFile(t *testing.T) {
	var filename = "test_data/empty"
	output, err := builtin.Wc.Run([]command.CommandArgument{command.CommandArgument(filename)}, nil)
	if err != nil || output.ExitCode != 0 {
		t.Fatalf("non-zero exitcode")
	}
	parseAndAssertOutput(t, output, true, true, true, true, filename, 0, 0, 0)
}

func TestWcFile(t *testing.T) {
	var filename = "test_data/advanced"
	output, err := builtin.Wc.Run([]command.CommandArgument{command.CommandArgument(filename)}, nil)
	if err != nil || output.ExitCode != 0 {
		t.Fatalf("non-zero exitcode")
	}
	parseAndAssertOutput(t, output, true, true, true, true, filename, 7, 14, 108)
}
