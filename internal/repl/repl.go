package repl

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/eqimd/bashgo/internal/bash"
	"github.com/eqimd/bashgo/internal/command/builtin"
)

/*
 * Структура для Read-Evaluate-Print Loop
 */
type Repl struct {
	bash bash.Bash
}

func (repl *Repl) StartRepl() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("repl> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error while reading from stdin: %w", err)
		}

		// convert CRLF to LF

		cmd = strings.Replace(cmd, "\n", "", -1)

		outp, exitCode, err := repl.bash.Execute(cmd)
		if err != nil {
			if errors.Is(err, builtin.ErrExit) {
				return nil
			}
			fmt.Println(err.Error())
		} else {
			fmt.Println(outp)
			if exitCode != 0 {
				fmt.Println("Non-zero exit code:", exitCode)
			}
		}
	}
}

func NewRepl(bash bash.Bash) *Repl {
	return &Repl{bash}
}
