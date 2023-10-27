package builtin

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
	"github.com/urfave/cli/v2"
)

type grepRunner struct{}

func (r *grepRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	var caseInsensitive bool
	var findWord bool
	var afterLines int
	var word string
	var filename string

	cliArgs := []string{"grep"}
	for _, arg := range args {
		cliArgs = append(cliArgs, string(arg))
	}

	cliApp := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "insensitive", Aliases: []string{"i"}},
			&cli.BoolFlag{Name: "word", Aliases: []string{"w"}},
			&cli.IntFlag{Name: "After", Aliases: []string{"A"}},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() < 1 {
				return fmt.Errorf("not enough arguments")
			}

			word = cCtx.Args().Get(0)

			if cCtx.Args().Len() > 1 {
				filename = cCtx.Args().Get(1)
			}

			caseInsensitive = cCtx.Bool("insensitive")
			findWord = cCtx.Bool("word")
			afterLines = cCtx.Int("After")

			return nil
		},
	}

	if err := cliApp.Run(cliArgs); err != nil {
		return nil, err
	}

	var grepData string
	if filename == "" {
		if input == nil {
			return nil, fmt.Errorf("input is nil")
		}
		grepData = input.Data
	} else {
		b, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("can't read file: %v", err)
		}

		grepData = string(b)
	}

	finalData, err := doGrep(caseInsensitive, findWord, afterLines, word, grepData)
	if err != nil {
		return nil, err
	}

	var builder strings.Builder
	for _, s := range finalData {
		builder.WriteString(s)
		builder.WriteString(" grep\n")
	}

	return &data.Output{Data: builder.String()}, nil
}

func doGrep(
	caseInsensitive bool,
	findWord bool,
	afterLines int,
	word string,
	data string,
) ([]string, error) {
	outp := []string{}
	splitData := strings.Split(data, "\n")

	if caseInsensitive {
		word = strings.ToLower(word)
		data = strings.ToLower(data)
	}

	shouldPrintAfterLines := 0
	if findWord {
		for _, s := range splitData {
			if s == word {
				outp = append(outp, s)
				shouldPrintAfterLines = afterLines
				continue
			}

			if shouldPrintAfterLines != 0 {
				outp = append(outp, s)
				shouldPrintAfterLines--
			}
		}

		return splitData, nil
	}

	rgx, err := regexp.Compile()
	if err != nil {
		return nil, err
	}
}
