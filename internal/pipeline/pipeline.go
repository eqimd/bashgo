package pipeline

import "github.com/eqimd/bashgo/internal/data"

/*
 * Интерфейс для пайплайна. Умеет запускаться
 */
type Pipeline interface {
	Run(input *data.Input) (*data.Output, error)
}
