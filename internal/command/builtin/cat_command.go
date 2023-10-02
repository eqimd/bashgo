package builtin

import (
	"io"
	"os"
	"strings"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

type catRunner struct{}

func (r *catRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	var contents = make([]string, 0, len(args))
	for _, arg := range args {
		filename := string(arg)
		file, err := os.Open(filename)
		if err != nil {
			return &data.Output{ExitCode: 1}, err
		}
		buf, err := io.ReadAll(file)
		if err != nil {
			return &data.Output{ExitCode: 255}, err
		}
		contents = append(contents, string(buf))
	}
	return &data.Output{Data: strings.Join(contents, "\n"), ExitCode: 0}, nil
}
