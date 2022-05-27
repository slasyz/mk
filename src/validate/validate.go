package validate

import (
	"fmt"

	"github.com/slasyz/mk/src/schema"
)

func validateParams(params []schema.Param) error {
	for i, param := range params {
		if param.Name == "" {
			return fmt.Errorf("empty param #%d name", i+1)
		}
	}

	return nil
}

func validate(commands []schema.Command, args string) error {
	set := make(map[string]struct{}, len(commands))
	for _, command := range commands {
		thisCmd := args + " " + command.Name
		if _, ok := set[command.Name]; ok {
			return fmt.Errorf("duplicated \"%s\" command", command.Name)
		}
		set[command.Name] = struct{}{}

		// Checking that there are no required params after optional one
		onlyOptionals := false
		onlyOptionalsName := ""
		for _, param := range command.Params {
			if param.Optional {
				if onlyOptionalsName == "" {
					onlyOptionalsName = param.Name
				}
				onlyOptionals = true
			} else if onlyOptionals {
				return fmt.Errorf("command \"%s\": all params after \"%s\" must be optional", thisCmd, onlyOptionalsName)
			}
		}

		err := validateParams(command.Params)
		if err != nil {
			return fmt.Errorf("error validating params: %w", err)
		}

		err = validate(command.Subcommands, thisCmd)
		if err != nil {
			return fmt.Errorf("error getting subcommand help: %w", err)
		}
	}

	return nil
}

func Validate(r *schema.Root) error {
	return validate(r.Commands, "mk")
}
