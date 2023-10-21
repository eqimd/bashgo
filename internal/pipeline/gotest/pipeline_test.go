package gotest

import (
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/external"
	"github.com/eqimd/bashgo/internal/pipe"
	"github.com/eqimd/bashgo/internal/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestPipelineImpl(t *testing.T) {
	pipes := []*pipe.Pipe{
		pipe.NewPipe(
			external.NewExternalCommand("echo"),
			[]command.CommandArgument{command.CommandArgument("pipeline test")},
		),
		pipe.NewPipe(
			external.NewExternalCommand("wc"),
			[]command.CommandArgument{},
		),
	}

	pipeline := pipeline.NewPipelineImpl(pipes)
	outp, err := pipeline.Run(nil)

	assert.NoError(t, err, "no error expected")
	assert.Equal(t, "      1       2      14\n", outp.Data)
}
