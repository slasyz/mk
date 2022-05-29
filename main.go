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
)

func maybePrintHelp(root *schema.Node, err error) {
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
	root, err := schema.Load("mk.yml")
	if err != nil {
		return fmt.Errorf("error parsing mk.yml: %w", err)
	}

	complete.Complete(root)

	lggr := logger.New(os.Stderr)
	shll := shell.New(
		"/bin/bash",
		shell.WithLogger(lggr),
	)

	script, err := build.Build(root, os.Args[1:])
	maybePrintHelp(root, err)
	if err != nil {
		return fmt.Errorf("error building script: %w", err)
	}

	err = shll.Exec(script.Cmd, script.Args, script.Workdir)
	maybePrintHelp(root, err)
	return err
}

func main() {
	err := _main()
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}
