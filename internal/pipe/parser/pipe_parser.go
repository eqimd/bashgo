package parser

import "github.com/eqimd/bashgo/internal/pipe"

type PipeParser interface {
	Parse(s string) (*pipe.Pipe, error)
}
