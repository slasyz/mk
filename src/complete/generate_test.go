package complete

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/slasyz/mk/src/schema"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		compLine  string
		compPoint int
		compWord  int
		expected  []string
	}{
		{
			"mk ",
			3,
			1,
			[]string{"noargs", "multiline", "subcommands", "params-arbitrary", "params-documented", "subcmd-params"},
		},
		{
			"mk noa",
			6,
			1,
			[]string{"noargs", "multiline", "subcommands", "params-arbitrary", "params-documented", "subcmd-params"},
		},
		{
			"mk noargs",
			6,
			1,
			[]string{"noargs", "multiline", "subcommands", "params-arbitrary", "params-documented", "subcmd-params"},
		},
		{
			"mk noargs",
			9,
			1,
			[]string{"noargs", "multiline", "subcommands", "params-arbitrary", "params-documented", "subcmd-params"},
		},
		{
			"mk noargs ",
			10,
			2,
			nil,
		},
		{
			"mk unexpected ",
			14,
			2,
			nil,
		},
		{
			"mk params-documented ",
			21,
			2,
			[]string{"true", "false"},
		},
		{
			"mk params-documented true ",
			26,
			3,
			[]string{"value1", "value2", "value3", "value4"},
		},
		{
			"mk subcmd-params ",
			17,
			2,
			[]string{"true", "false", "subfirst", "subsecond"},
		},
	}

	root, err := schema.Parse("../../examples/example.yml")
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s/%d/%d", tt.compLine, tt.compWord, tt.compPoint), func(t *testing.T) {
			res, err := Generate(root, tt.compLine, tt.compPoint, tt.compWord)
			require.NoError(t, err)
			require.ElementsMatch(t, tt.expected, res)
		})
	}
}
