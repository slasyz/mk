package build

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/slasyz/mk/src/schema"
)

func TestBuildScript(t *testing.T) {
	tests := []struct {
		name           string
		filename       string
		args           []string
		expectedScript *Script
		expectedError  string
	}{
		{
			name:     "noargs",
			filename: "../../examples/example.yml",
			args:     []string{"noargs"},
			expectedScript: &Script{
				Cmd:  `echo "It works!"`,
				Args: []string{},
			},
			expectedError: "",
		},
		{
			name:     "noargs with args",
			filename: "../../examples/example.yml",
			args:     []string{"noargs", "unexpected"},
			expectedScript: &Script{
				Cmd:  `echo "It works!"`,
				Args: []string{"unexpected"},
			},
			expectedError: "",
		},
		{
			name:     "subcommand root",
			filename: "../../examples/example.yml",
			args:     []string{"subcommands"},
			expectedScript: &Script{
				Cmd:  `echo "Root command."`,
				Args: []string{},
			},
			expectedError: "",
		},
		{
			name:     "subcommand first",
			filename: "../../examples/example.yml",
			args:     []string{"subcommands", "first"},
			expectedScript: &Script{
				Cmd:  `echo "Args for first subcommand are $@"`,
				Args: []string{},
			},
			expectedError: "",
		},
		{
			name:     "subcommand second",
			filename: "../../examples/example.yml",
			args:     []string{"subcommands", "second", "arg1234"},
			expectedScript: &Script{
				Cmd:  `echo "Args for second subcommand are $@"`,
				Args: []string{"arg1234"},
			},
			expectedError: "",
		},
		{
			name:     "subcommand unexpected",
			filename: "../../examples/example.yml",
			args:     []string{"subcommands", "unexpected"},
			expectedScript: &Script{
				Cmd:  `echo "Root command."`,
				Args: []string{"unexpected"},
			},
			expectedError: "",
		},
		{
			name:     "params documented valid",
			filename: "../../examples/example.yml",
			args:     []string{"params-documented", "true", "value3", "123"},
			expectedScript: &Script{
				Cmd:  `echo "Documented parameters list is $@"`,
				Args: []string{"true", "value3", "123"},
			},
			expectedError: "",
		},
		{
			name:           "params documented invalid",
			filename:       "../../examples/example.yml",
			args:           []string{"params-documented", "true", "invalid", "123"},
			expectedScript: nil,
			expectedError:  "error validating params: param tag must be one of these: value1, value2, value3, value4, but it's invalid",
		},
		{
			name:     "params documented omit optional",
			filename: "../../examples/example.yml",
			args:     []string{"params-documented", "true", "value1"},
			expectedScript: &Script{
				Cmd:  `echo "Documented parameters list is $@"`,
				Args: []string{"true", "value1"},
			},
			expectedError: "",
		},
		{
			name:           "params documented extra",
			filename:       "../../examples/example.yml",
			args:           []string{"params-documented", "true", "value2", "123", "321"},
			expectedScript: nil,
			expectedError:  "error validating params: unexpected argument \"321\"",
		},
		{
			name:     "include file",
			filename: "../../examples/example.yml",
			args:     []string{"inner", "included", "param1"},
			expectedScript: &Script{
				Cmd:  `echo "Inner command."`,
				Args: []string{"param1"},
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, err := schema.Load(tt.filename)
			require.NoError(t, err)

			res, err := Build(root, tt.args)
			if tt.expectedError != "" {
				require.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedScript, res)
			}
		})
	}
}
