package complete

import (
	"fmt"

	"github.com/google/shlex"

	"github.com/slasyz/mk/src/schema"
)

//func searchSubcommand(name string, commands []schema.Command) (schema.Command, bool) {
//return
//}

func Generate(root *schema.Node, compLine string, compPoint, compWord int) ([]string, error) {
	cmds, err := shlex.Split(compLine)
	if err != nil {
		return nil, fmt.Errorf("error splitting comp line: %w", err)
	}

	// TODO: rewrite this

	commands := root.Children
	var paramValues []string
	var params []schema.Param
	var i int
	for i = 1; i < compWord; i++ {
		searchingFor := cmds[i]

		movedDeeper := false
		for _, cmd := range commands {
			if cmd.Name == searchingFor {
				commands = cmd.Children
				params = cmd.Params
				movedDeeper = true
				break
			}
		}
		if !movedDeeper {
			commands = nil
			break
		}
	}

	shift := compWord - i
	if len(params) > shift {
		paramValues = params[shift].Values
	}

	var res []string
	for _, cmd := range commands {
		res = append(res, cmd.Name)
	}
	for _, paramValue := range paramValues {
		res = append(res, paramValue)
	}

	return res, nil
}
