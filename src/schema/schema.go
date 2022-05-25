package schema

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Root struct {
	Version  string    `yaml:"version"`
	Commands []Command `json:"commands"`
}

type Command struct {
	Name        string    `yaml:"name"`
	Help        string    `yaml:"help"`
	Params      []Param   `yaml:"params"`
	Cmd         string    `yaml:"cmd"`
	Subcommands []Command `yaml:"subcommands"`
}

type Param struct {
	Name     string   `yaml:"name"`
	Values   []string `yaml:"values"`
	Optional bool     `yaml:"optional"`
}

func Parse(filename string) (*Root, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	var root Root
	err = yaml.NewDecoder(f).Decode(&root)
	if err != nil {
		return nil, fmt.Errorf("error decoding yaml: %w", err)
	}

	return &root, nil
}
