package parser

import "github.com/eqimd/bashgo/internal/pipe"

type PipeParser interface {
	// Метод создаёт исполяемый Pipe из строки, описывающей его
	Parse(s string) (*pipe.Pipe, error)
}
