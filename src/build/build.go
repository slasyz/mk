package build

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/slasyz/mk/src/schema"
)

var ErrInvalidCommand = errors.New("invalid command")

// Script is a structure containing all information for runner to execute the target.
// It's script text, arguments to pass to it, maybe environment variables.
// Other scripts to run before it.
type Script struct {
	Cmd  string
	Args []string
}

func validateParamsArgs(cmd string, params []schema.Param, args []string) error {
	if cmd == "" {
		return errors.New("empty command")
	}
	if len(params) == 0 {
		return nil
	}

	if len(args) > len(params) {
		return fmt.Errorf("unexpected argument \"%s\"", args[len(params)])
	}
	if len(args) < len(params) && !params[len(args)].Optional {
		return fmt.Errorf("param %s is required", params[len(args)].Name)
	}
	for i, arg := range args {
		param := params[i]
		if len(param.Values) > 0 && !slices.Contains(param.Values, arg) {
			return fmt.Errorf("param %s must be one of these: %s, but it's %s", param.Name, strings.Join(param.Values, ", "), arg)
		}
	}

	return nil
}

func scriptThis(command *schema.Command, args []string) (*Script, error) {
	err := validateParamsArgs(command.Cmd, command.Params, args)
	if err != nil {
		return nil, fmt.Errorf("error validating params: %w", err)
	}

	return &Script{
		Cmd:  command.Cmd,
		Args: args,
	}, nil
}

func scriptThisOrDeeper(command *schema.Command, args []string) (*Script, error) {
	if len(command.Subcommands) > 0 && len(args) > 0 {
		return scriptSelectCommand(command, command.Subcommands, args)
	}

	return scriptThis(command, args)
}

func scriptSelectCommand(command *schema.Command, subcommands []schema.Command, args []string) (*Script, error) {
	for _, subcommand := range subcommands {
		if subcommand.Name == args[0] {
			return scriptThisOrDeeper(&subcommand, args[1:])
		}
	}

	if command == nil {
		return nil, ErrInvalidCommand
	}

	return scriptThis(command, args)
}

func Build(root *schema.Root, args []string) (*Script, error) {
	if len(args) == 0 || args[0] == "help" {
		return nil, ErrInvalidCommand
	}

	return scriptSelectCommand(nil, root.Commands, args)
}
