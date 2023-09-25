package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eqimd/bashgo/internal/bash"
)

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
		// TODO
		fmt.Println(outp, exitCode, err)
	}
}

func NewRepl(bash bash.Bash) *Repl {
	return &Repl{bash}
}
