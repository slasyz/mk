package help

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/slasyz/mk/src/schema"
)

type example struct {
	help string
	cmd  string
}

func buildParamsStr(params []schema.Param) (string, error) {
	var cmdParams []string
	for i, param := range params {
		cmdParam := ""
		if len(param.Values) > 0 {
			cmdParam = "{" + strings.Join(param.Values, "|") + "}"
		} else if param.Name != "" {
			cmdParam = strings.ToUpper(param.Name)
		} else {
			return "", fmt.Errorf("empty param #%d", i+1)
		}

		if param.Optional {
			cmdParam = "[" + cmdParam + "]"
		}

		cmdParams = append(cmdParams, cmdParam)
	}

	return strings.Join(cmdParams, " "), nil
}

func getExamples(commands []*schema.Node, args string) ([]example, error) {
	var res []example

	for _, command := range commands {
		thisCmd := args + " " + command.Name

		if command.Cmd != "" {
			thisExample := example{
				help: command.Help,
				cmd:  thisCmd,
			}

			paramsStr, err := buildParamsStr(command.Params)
			if err != nil {
				return nil, fmt.Errorf("error building params for %s command", thisCmd)
			}
			if paramsStr != "" {
				thisExample.cmd += " " + paramsStr
			}

			res = append(res, thisExample)
		}

		resDeeper, err := getExamples(command.Children, thisCmd)
		if err != nil {
			return nil, fmt.Errorf("error getting subcommand help: %w", err)
		}

		res = append(res, resDeeper...)
	}

	return res, nil
}

func Help(root *schema.Node) (string, error) {
	b := bytes.NewBuffer(nil)

	examples, err := getExamples(root.Children, "mk")
	if err != nil {
		return "", err
	}

	for _, ex := range examples {
		fmt.Fprintln(b, ex.help)
		fmt.Fprintf(b, "  $ %s\n", ex.cmd)
		fmt.Fprintln(b)
	}

	return b.String(), nil
}
