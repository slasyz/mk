package schema

type MkFile struct {
	Version  string    `yaml:"version"`
	Commands []Command `json:"commands"`
}

type Command struct {
	Name string `yaml:"name"`
	Help string `yaml:"help"`

	Cmd     string  `yaml:"cmd"`
	Workdir string  `yaml:"workdir"`
	Params  []Param `yaml:"params"`

	Subcommands []Command `yaml:"subcommands"`
	Include     string    `yaml:"include"`
}

type Param struct {
	Name     string   `yaml:"name"`
	Values   []string `yaml:"values"`
	Optional bool     `yaml:"optional"`
}
