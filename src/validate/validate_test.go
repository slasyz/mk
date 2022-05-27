package validate

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/slasyz/mk/src/schema"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		filename    string
		expectedErr string
	}{
		{
			filename:    "invalid_optionals.yml",
			expectedErr: `command "mk required-after-optional": all params after "tag" must be optional`,
		},
		{
			filename:    "invalid_unnamed.yml",
			expectedErr: "error validating params: empty param #3 name",
		},
	}

	for _, tt := range tests {
		t.Run(strings.TrimSuffix(tt.filename, ".yml"), func(t *testing.T) {
			root, err := schema.Parse(filepath.Join("..", "..", "examples", tt.filename))
			require.NoError(t, err)

			err = Validate(root)
			if tt.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.expectedErr)
			}
		})
	}
}
