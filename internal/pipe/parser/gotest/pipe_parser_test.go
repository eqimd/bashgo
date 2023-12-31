package parser_test

import (
	"os"
	"testing"

	"github.com/eqimd/bashgo/internal/command/splitter"
	"github.com/eqimd/bashgo/internal/pipe/parser"
	"github.com/stretchr/testify/require"
)

// Данный тест проверяет корректность построения Pipe,
// Корректность Split проверяется в отдельных тестах
func TestPipeParserSimple(t *testing.T) {
	text := "abcdabcd"
	line := "echo " + text
	splitter := splitter.NewCommandSplitterImpl()
	parser := parser.NewPipeParserImpl(splitter)

	pipe := parser.Parse(line)

	out, err := pipe.RunPipe(nil)

	require.NoError(t, err, "unexpected error: %v")
	require.Equal(t, text, out.Data)
}

// Тест проверяет, что парсер успешно выполняет подстановку одной переменной среды окружения
func TestPipeParserWithEnvVariable(t *testing.T) {
	text := "text"
	err := os.Setenv("test", text)
	require.NoError(t, err)

	line := "echo $test"

	splitter := splitter.NewCommandSplitterImpl()
	parser := parser.NewPipeParserImpl(splitter)

	pipe := parser.Parse(line)

	out, err := pipe.RunPipe(nil)

	require.NoError(t, err)
	require.Equal(t, text, out.Data)
}

// Тест проверяет, что парсер успешно выполняет подстановку нескольких переменных среды окружения, в том числе в качестве имени команды
func TestPipeParserWithMultipleEnvVariables(t *testing.T) {
	text := "text"
	err := os.Setenv("com1", "ec")
	require.NoError(t, err)
	err = os.Setenv("com2", "ho")
	require.NoError(t, err)
	err = os.Setenv("test", text)
	require.NoError(t, err)

	line := "$com1$com2 $test/"

	splitter := splitter.NewCommandSplitterImpl()
	parser := parser.NewPipeParserImpl(splitter)

	pipe := parser.Parse(line)

	out, err := pipe.RunPipe(nil)

	require.NoError(t, err)
	require.Equal(t, text+"/", out.Data)
}

// Тест проверяет, что парсер успешно обрабатывает подстановку переменной окружения в одинарных кавычках
func TestPipeParserWithEnvVariableInSingleQuotes(t *testing.T) {
	text := "text"
	err := os.Setenv("test", text)
	require.NoError(t, err)

	line := "echo '$test'"

	splitter := splitter.NewCommandSplitterImpl()
	parser := parser.NewPipeParserImpl(splitter)

	pipe := parser.Parse(line)

	out, err := pipe.RunPipe(nil)

	require.NoError(t, err)
	require.Equal(t, "$test", out.Data)
}

// Тест проверяет, что парсер успешно обрабатывает подстановку переменной окружения в двойных кавычках
func TestPipeParserWithEnvVariableInDoubleQuotes(t *testing.T) {
	text := "text"
	err := os.Setenv("test", text)
	require.NoError(t, err)

	line := "echo \"$test\""

	splitter := splitter.NewCommandSplitterImpl()
	parser := parser.NewPipeParserImpl(splitter)

	pipe := parser.Parse(line)

	out, err := pipe.RunPipe(nil)

	require.NoError(t, err)
	require.Equal(t, text, out.Data)
}
