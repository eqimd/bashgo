package parser

import (
	"os"
	"unicode"

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
	curQuotes := rune(' ')
	envVarStartIndex := -1
	newS := ""
	s = s + " "
	for i, ch := range s {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			if envVarStartIndex != -1 {
				// закончилось имя переменной
				envName := s[envVarStartIndex+1 : i]
				envValue := os.Getenv(envName)
				newS = newS + envValue

				envVarStartIndex = -1
			}
		}
		if ch == rune('$') && curQuotes != rune('\'') {
			envVarStartIndex = i
		}

		if envVarStartIndex == -1 {
			newS = newS + string(ch)
		}

		if curQuotes == rune(' ') && (ch == rune('\'') || ch == rune('"')) {
			curQuotes = ch
		} else if curQuotes == ch {
			curQuotes = rune(' ')
		}
	}

	command, args := parser.splitter.Split(newS)

	return pipe.NewPipe(command, args)
}

// Создание нового парсера
func NewPipeParserImpl(splitter splitter.CommandSplitter) *PipeParserImpl {
	return &PipeParserImpl{splitter}
}
