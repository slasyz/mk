package schema

import (
	"path/filepath"
	"strings"
	"testing"

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
			expectedErr: `duplicated "cmdname" command`,
		},
		{
			filename:    "invalid_optionals.yml",
			expectedErr: `error validating params for "mk required-after-optional": all params after "tag" must be optional`,
		},
		{
			filename:    "invalid_unnamed.yml",
			expectedErr: `error validating params for "mk unnamed": empty param #3 name`,
		},
	}

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
