package gotest

import (
	"testing"

	"github.com/eqimd/bashgo/internal/pipeline/splitter"
	"github.com/stretchr/testify/assert"
)

func TestPipelineSplitterImpl_NoPipes(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo test"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{cmd}, splitCmds)
}

func TestPipelineSplitterImpl_OnePipe(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo test | echo test2"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo test", "echo test2"}, splitCmds)
}

func TestPipelineSplitterImpl_OnePipe_NoSpaces(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo test|echo test2"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo test", "echo test2"}, splitCmds)
}

func TestPipelineSplitterImpl_MultiPipes(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo test|echo test2 | echo abra|echo notabra |   echo echo       | notecho"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo test", "echo test2", "echo abra", "echo notabra", "echo echo", "notecho"}, splitCmds)
}

func TestPipelineSplitterImpl_SingleQuotes_NoPipe(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo 'test'"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo 'test'"}, splitCmds)
}

func TestPipelineSplitterImpl_SingleQuotes_OnePipe(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo 'test 1' | echo test2"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo 'test 1'", "echo test2"}, splitCmds)
}

func TestPipelineSplitterImpl_SingleQuotes_MultiPipe(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo 'test' | echo test2    |echo 'test3'|echo 'test4'|echo 'test 5'|echo test6"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo 'test'", "echo test2", "echo 'test3'", "echo 'test4'", "echo 'test 5'", "echo test6"}, splitCmds)
}

func TestPipelineSplitterImpl_DoubleQuotes(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo \"test\" | echo test2    |echo \"test3\"|echo \"test4\"|echo \"test 5\"|echo test6"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo \"test\"", "echo test2", "echo \"test3\"", "echo \"test4\"", "echo \"test 5\"", "echo test6"}, splitCmds)
}

func TestPipelineSplitterImpl_AllQuotes(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo \"test | 'inside' \" | echo test2    |echo 'test3 | \"inside\" '|echo 'test4'|echo \"test 5\"|echo test6"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo \"test | 'inside' \"", "echo test2", "echo 'test3 | \"inside\" '", "echo 'test4'", "echo \"test 5\"", "echo test6"}, splitCmds)
}

func TestPipelineSplitterImpl_LastSymbolPipe(t *testing.T) {
	splitter := splitter.NewPipelineSplitterImpl()

	cmd := "echo blin |"
	splitCmds := splitter.Split(cmd)

	assert.Equal(t, []string{"echo blin", ""}, splitCmds)
}
