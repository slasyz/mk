package schema

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Node struct {
	Name   string
	Help   string
	Params []Param
	Cmd    string

	Children []*Node
}

func validateParams(params []Param) error {
	onlyOptionals := false
	onlyOptionalsName := ""
	for i, param := range params {
		// Checking that there are no required params after optional one
		if param.Optional {
			if onlyOptionalsName == "" {
				onlyOptionalsName = param.Name
			}
			onlyOptionals = true
		} else if onlyOptionals {
			return fmt.Errorf("all params after \"%s\" must be optional", onlyOptionalsName)
		}

		if param.Name == "" {
			return fmt.Errorf("empty param #%d name", i+1)
		}
	}

	return nil
}

func commandsToNodes(commands []Command, path []string) ([]*Node, error) {
	if len(commands) == 0 {
		return nil, nil
	}

	res := make([]*Node, len(commands))
	namesSet := make(map[string]struct{}, len(commands))
	for i, command := range commands {
		cmdPath := append(path, command.Name)

		if _, ok := namesSet[command.Name]; ok {
			return nil, fmt.Errorf("duplicated \"%s\" command", command.Name)
		}
		namesSet[command.Name] = struct{}{}

		err := validateParams(command.Params)
		if err != nil {
			return nil, fmt.Errorf(`error validating params for "%s": %w`, strings.Join(cmdPath, " "), err)
		}

		children, err := commandsToNodes(command.Subcommands, cmdPath)
		if err != nil {
			return nil, fmt.Errorf(`error converting "%s" children to nodes: %w`, strings.Join(cmdPath, " "), err)
		}

		res[i] = &Node{
			Name:     command.Name,
			Help:     command.Help,
			Params:   command.Params,
			Cmd:      command.Cmd,
			Children: children,
		}
	}

	return res, nil
}

func toTree(mkFile *MkFile) (*Node, error) {
	children, err := commandsToNodes(mkFile.Commands, []string{"mk"})
	if err != nil {
		return nil, err
	}

	return &Node{
		Name:     "mk",
		Children: children,
	}, nil
}

func Load(filename string) (*Node, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	var mkfile MkFile
	err = yaml.NewDecoder(f).Decode(&mkfile)
	if err != nil {
		return nil, fmt.Errorf("error decoding yaml: %w", err)
	}

	return toTree(&mkfile)
}
