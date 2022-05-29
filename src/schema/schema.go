package schema

type MkFile struct {
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
