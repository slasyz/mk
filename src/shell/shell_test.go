package shell

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShell_Exec(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	shell := New(
		"/bin/bash",
		WithStdout(&stdout),
		WithStderr(&stderr),
	)

	err := shell.Exec(`
		echo "Hey, $1! I'm $2."
		echo "stderr example" >&2
`, []string{"Slava", "Petya"}, "")
	require.NoError(t, err)

	assert.Equal(t, "Hey, Slava! I'm Petya.\n", stdout.String())
	assert.Equal(t, "stderr example\n", stderr.String())
}

func TestShell_Noargs(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	shell := New(
		"/bin/bash",
		WithStdout(&stdout),
		WithStderr(&stderr),
	)

	err := shell.Exec(`echo "Hey, Slava!"`, nil, "")
	require.NoError(t, err)

	assert.Equal(t, "Hey, Slava!\n", stdout.String())
	assert.Empty(t, stderr.String())
}

func TestShell_Error(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	shell := New(
		"/bin/bash",
		WithStdout(&stdout),
		WithStderr(&stderr),
	)

	err := shell.Exec(`exit 1`, nil, "")
	require.Error(t, err)
}

func TestShell_Workdir(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	shell := New(
		"/bin/bash",
		WithStdout(&stdout),
		WithStderr(&stderr),
	)

	err := shell.Exec(`pwd`, nil, "/")
	require.NoError(t, err)

	require.Equal(t, "/\n", stdout.String())
}
