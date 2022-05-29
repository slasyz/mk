package schema

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	errIncludeAndCmd         = errors.New("command cannot include file and contain command")
	errIncludeAndSubcommands = errors.New("command cannot include file and contain subcommands")
	errIncludeAndParams      = errors.New("command cannot include file and contain params")
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

func commandToNode(currentFile string, command Command, path []string) (*Node, error) {
	node := Node{
		Name: command.Name,
		Help: command.Help,
	}

	if command.Include != "" {
		if command.Cmd != "" {
			return nil, errIncludeAndCmd
		}
		if len(command.Subcommands) > 0 {
			return nil, errIncludeAndSubcommands
		}
		if len(command.Params) > 0 {
			return nil, errIncludeAndParams
		}

		includeNode, err := Load(filepath.Join(filepath.Dir(currentFile), command.Include))
		if err != nil {
			return nil, fmt.Errorf("error loading %s: %w", command.Include, err)
		}

		node.Children = includeNode.Children
		return &node, nil
	}

	node.Cmd = command.Cmd

	err := validateParams(command.Params)
	if err != nil {
		return nil, fmt.Errorf(`validating params: %w`, err)
	}
	node.Params = command.Params

	node.Children, err = commandsToNodes(currentFile, command.Subcommands, path)
	if err != nil {
		return nil, fmt.Errorf(`converting subcommands to tree nodes: %w`, err)
	}

	return &node, nil
}

func commandsToNodes(currentFile string, commands []Command, path []string) ([]*Node, error) {
	if len(commands) == 0 {
		return nil, nil
	}

	res := make([]*Node, len(commands))
	namesSet := make(map[string]struct{}, len(commands))
	for i, command := range commands {
		cmdPath := append(path, command.Name)

		if _, ok := namesSet[command.Name]; ok {
			return nil, fmt.Errorf("duplicated \"%s\" command", strings.Join(cmdPath, " "))
		}
		namesSet[command.Name] = struct{}{}

		node, err := commandToNode(currentFile, command, cmdPath)
		if err != nil {
			return nil, fmt.Errorf("error in \"%s\": %w", strings.Join(cmdPath, " "), err)
		}

		res[i] = node
	}

	return res, nil
}

func toTree(currentFile string, mkFile *MkFile) (*Node, error) {
	children, err := commandsToNodes(currentFile, mkFile.Commands, []string{"mk"})
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

	return toTree(filename, &mkfile)
}
