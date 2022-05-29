package schema

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad_Basic(t *testing.T) {
	tests := []struct {
		filename    string
		expectedErr string
	}{
		{
			filename:    "example.yml",
			expectedErr: "",
		},
		{
			filename:    "example_inner.yml",
			expectedErr: "",
		},
		{
			filename:    "invalid_duplicate.yml",
			expectedErr: `duplicated "mk cmdname" command`,
		},
		{
			filename:    "invalid_include_params.yml",
			expectedErr: `error in "mk include_params_conflict": command cannot include file and contain params`,
		},
		{
			filename:    "invalid_include_cmd.yml",
			expectedErr: `error in "mk include_cmd_conflict": command cannot include file and contain command`,
		},
		{
			filename:    "invalid_include_scmds.yml",
			expectedErr: `error in "mk include_subcommands_conflict": command cannot include file and contain subcommands`,
		},
		{
			filename:    "invalid_optionals.yml",
			expectedErr: `error in "mk required-after-optional": validating params: all params after "tag" must be optional`,
		},
		{
			filename:    "invalid_unnamed.yml",
			expectedErr: `error in "mk unnamed": validating params: empty param #3 name`,
		},
	}

	// Just checking if I didn't forget to test all examples
	total, err := filepath.Glob("../../examples/*.yml")
	require.NoError(t, err)
	todo, err := filepath.Glob("../../examples/todo_*.yml")
	require.NoError(t, err)
	assert.Equal(t, len(tests), len(total)-len(todo))

	for _, tt := range tests {
		t.Run(strings.TrimSuffix(tt.filename, ".yml"), func(t *testing.T) {
			_, err := Load(filepath.Join("..", "..", "examples", tt.filename))
			if tt.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.expectedErr)
			}
		})
	}
}
