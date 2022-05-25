package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/slasyz/mk/src/build"
	"github.com/slasyz/mk/src/complete"
	"github.com/slasyz/mk/src/help"
	"github.com/slasyz/mk/src/logger"
	"github.com/slasyz/mk/src/schema"
	"github.com/slasyz/mk/src/shell"
	"github.com/slasyz/mk/src/validate"
)

func maybePrintHelp(root *schema.Root, err error) {
	if errors.Is(err, build.ErrInvalidCommand) {
		helpText, err := help.Help(root)
		if err != nil {
			fmt.Println("Error generating help: %w", err)
		}

		fmt.Print(helpText)
		os.Exit(1)
	}
}

func _main() error {
	root, err := schema.Parse("mk.yml")
	if err != nil {
		return fmt.Errorf("error parsing mk.yml: %w", err)
	}

	lggr := logger.New(os.Stderr)
	shll := shell.New(
		"/bin/bash",
		shell.WithLogger(lggr),
	)

	err = validate.Validate(root)
	if err != nil {
		return fmt.Errorf("invalid mk.yml: %w", err)
	}

	script, err := build.Build(root, os.Args[1:])
	maybePrintHelp(root, err)
	if err != nil {
		return fmt.Errorf("error building script: %w", err)
	}

	err = shll.Exec(script.Cmd, script.Args)
	maybePrintHelp(root, err)
	return err
}

func main() {
	complete.Complete()

	err := _main()
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}
