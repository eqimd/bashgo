package builtin

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)

var (
	Cat  = &builtinCommand{"cat", &catRunner{}}
	Echo = &builtinCommand{"echo", &echoRunner{}}
	Wc   = &builtinCommand{"wc", &wcRunner{}}
)

var ErrNoCommand = errors.New("command does not exist")

type builtinCommand struct {
	Name string

	command.Command
}

type catRunner struct{}

func (r *catRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}

type echoRunner struct{}

func (r *echoRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	// TODO
	return nil, nil
}

type wcRunner struct{}

func (r *wcRunner) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	filenames := make([]string, 0)
	needLines := false
	needWords := false
	needChars := false

	for _, arg := range args {
		var offset int
		strarg := string(arg)
		r, s := utf8.DecodeRuneInString(strarg)
		if r == '-' {
			for offset < len(strarg) {
				offset += s
				r, s = utf8.DecodeRuneInString(strarg[offset:])
				switch r {
				case 'l':
					needLines = true
				case 'w':
					needWords = true
				case 'c':
					needChars = true
				}
			}
		} else {
			filenames = append(filenames, strarg)
		}
	}
	if !(needLines || needWords || needChars) {
		needLines = true
		needWords = true
		needChars = true
	}

	if len(filenames) == 0 {
		inputReader := strings.NewReader(input.Data)
		result := make([]string, 0)
		err := calcAndAddWcStats(&result, needLines, needWords, needChars, inputReader)
		if err != nil {
			return &data.Output{ExitCode: 255}, err
		}
		return &data.Output{Data: strings.Join(result, " ")}, nil
	}

	results := make([]string, 0)
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return &data.Output{ExitCode: 1}, err
		}
		result := make([]string, 0)
		result = append(result, filename)
		err = calcAndAddWcStats(&result, needLines, needWords, needChars, file)
		if err != nil {
			return &data.Output{ExitCode: 255}, err
		}
		results = append(results, strings.Join(result, " "))
	}
	return &data.Output{Data: strings.Join(results, "\n")}, nil
}

func calcAndAddWcStats(resultHolder *[]string, needLines bool, needWords bool, needChars bool, reader io.Reader) error {
	bufReader := bufio.NewReader(reader)
	prevSeparator := false
	thisSeparator := false
	lines := 0
	words := 0
	chars := 0

	for {
		r, s, err := bufReader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		thisSeparator = r == ' ' || r == '\n' || r == '\t' || r == '\r'

		if r == '\n' {
			lines++
		}
		if !(prevSeparator) && thisSeparator {
			words++
		}
		chars += s

		prevSeparator = thisSeparator
	}
	if chars > 0 {
		lines++
	}
	if !(prevSeparator) && chars > 0 {
		words++
	}

	if needLines {
		*resultHolder = append(*resultHolder, fmt.Sprint(lines))
	}
	if needWords {
		*resultHolder = append(*resultHolder, fmt.Sprint(words))
	}
	if needChars {
		*resultHolder = append(*resultHolder, fmt.Sprint(chars))
	}

	return nil
}

func LookupBuiltinCommand(cmd string) (*builtinCommand, error) {
	switch cmd {
	case Cat.Name:
		return Cat, nil
	case Echo.Name:
		return Echo, nil
	case Wc.Name:
		return Wc, nil
	default:
		return nil, ErrNoCommand
	}
}
