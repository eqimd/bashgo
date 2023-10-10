package parser_test

import (
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

	pipe, err := parser.Parse(line)

	require.NoError(t, err, "unexpected error: %v")

	out, err := pipe.RunPipe(nil)

	require.NoError(t, err, "unexpected error: %v")
	require.Equal(t, text, out.Data)
}
