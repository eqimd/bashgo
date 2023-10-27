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
	originalData := data

	if caseInsensitive {
		word = strings.ToLower(word)
		data = strings.ToLower(data)
	}

	outp := []string{}

	splitData := strings.Split(data, "\n")
	splitData = splitData[:len(splitData)-1]

	splitOriginalData := strings.Split(originalData, "\n")
	splitOriginalData = splitOriginalData[:len(splitOriginalData)-1]

	var checkFunc (func(s string) bool)

	if findWord {
		checkFunc = func(s string) bool {
			splitS := strings.Split(s, " ")
			for _, sp := range splitS {
				if sp == s {
					return true
				}
			}

			return false
		}
	} else {
		rgx, err := regexp.Compile(word)
		if err != nil {
			return nil, err
		}

		checkFunc = func(s string) bool {
			return rgx.FindStringIndex(s) != nil
		}
	}

	shouldPrintAfterLines := 0
	for i, s := range splitData {
		if checkFunc(s) {
			outp = append(outp, splitOriginalData[i])
			shouldPrintAfterLines = afterLines
			continue
		}

		if shouldPrintAfterLines != 0 {
			outp = append(outp, splitOriginalData[i])
			shouldPrintAfterLines--
		}
	}

	return outp, nil
}
