package parser

import (
	"os"
	"regexp"
	"strings"

	"github.com/eqimd/bashgo/internal/command/splitter"
	"github.com/eqimd/bashgo/internal/pipe"
)

var _ PipeParser = &PipeParserImpl{}

// Структура для работы парсера, хранит внутри себя CommandSplitter
type PipeParserImpl struct {
	splitter splitter.CommandSplitter
}

// Метод создаёт исполяемый Pipe из строки, описывающей его, а также подставляет в требуемые места переменные окружения
func (parser *PipeParserImpl) Parse(s string) *pipe.Pipe {
	re := regexp.MustCompile(`\$\S*`)
	for {
		envWithD := re.FindString(s)
		envName, _ := strings.CutPrefix(envWithD, "$")
		if envName == "" {
			break
		}
		envValue := os.Getenv(envName)
		s = strings.ReplaceAll(s, envWithD, envValue)
	}

	command, args := parser.splitter.Split(s)

	return pipe.NewPipe(command, args)
}

// Создание нового парсера
func NewPipeParserImpl(splitter splitter.CommandSplitter) *PipeParserImpl {
	return &PipeParserImpl{splitter}
}
