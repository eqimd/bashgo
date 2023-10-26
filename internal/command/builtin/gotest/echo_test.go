package builtin_test

import (
	"strings"
	"testing"

	"github.com/eqimd/bashgo/internal/command"
	"github.com/eqimd/bashgo/internal/command/builtin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var Echo, _ = builtin.LookupBuiltinCommand("echo")

func TestEchoSingle(t *testing.T) {
	teststrs := []string{
		"",
		"abc",
		"abc\nabc",
		"\nabc",
		"\t1337$\t",
		"\x00\x01\x02\x03\x04\x05",
		"Привет!\nBonjour!\n¡Hola!\nمرحبا!\nこんにちは!",
	}
	for _, tstr := range teststrs {
		output, err := Echo.Run([]command.CommandArgument{command.CommandArgument(tstr)}, nil)

		require.NoError(t, err, "unexpected error: %v")
		require.Equal(t, 0, output.ExitCode, "wrong text echoed: want %v, got %v", 0, output.ExitCode)
		assert.Equal(t, tstr, output.Data, "wrong text echoed: want %v, got %v", tstr, output.Data)
	}
}

func TestEchoMultiple(t *testing.T) {
	teststrss := [][]string{
		{"abc", "abc"},
		{"abc", "def"},
		{"\x00", "\x01", "\x02", "\x03", "\x04", "\x05"},
		{"Привет!", "Bonjour!", "¡Hola!", "مرحبا!", "こんにちは!"},
	}
	for _, tstrs := range teststrss {
		args := make([]command.CommandArgument, 0, len(tstrs))
		exp := strings.Join(tstrs, " ")
		for _, tstr := range tstrs {
			args = append(args, command.CommandArgument(tstr))
		}
		output, err := Echo.Run(args, nil)

		require.NoError(t, err, "unexpected error: %v")
		require.Equal(t, 0, output.ExitCode, "wrong text echoed: want %v, got %v", 0, output.ExitCode)
		assert.Equal(t, exp, output.Data, "wrong text echoed: want %v, got %v", exp, output.Data)
	}
}
