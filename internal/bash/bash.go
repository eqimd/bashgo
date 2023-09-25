package bash

type Bash interface {
	Execute(command string) (string, int, error)
}
