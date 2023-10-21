package external

import (
	"errors"
	"io"
	"os/exec"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/data"
)


/*
 * Структура, описывающая внешнюю команду
 */
type ExternalCommand struct {
	executablePath string
}

func (cmd *ExternalCommand) Run(args []command.CommandArgument, input *data.Input) (*data.Output, error) {
	strArgs := make([]string, 0, len(args))
	for _, arg := range args {
		strArgs = append(strArgs, string(arg))
	}

	outp := &data.Output{}

	extCmd := exec.Command(cmd.executablePath, strArgs...)

	if input != nil {
		stdin, err := extCmd.StdinPipe()
		if err != nil {
			return nil, err
		}

		_, err = io.WriteString(stdin, input.Data)
		if err != nil {
			return nil, err
		}

		stdin.Close()
	}

	out, err := extCmd.CombinedOutput()

	outp.Data = string(out)

	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			outp.ExitCode = exitErr.ExitCode()
		} else {
			return nil, err
		}
	}

	return outp, nil
}

/*
 * Конструктор для внешней команды
 */
func NewExternalCommand(executablePath string) *ExternalCommand {
	return &ExternalCommand{executablePath}
}
