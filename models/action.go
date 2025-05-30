package models

type Action struct {
	Title string   `yaml:"title"`
	Cmd   string   `yaml:"cmd"`
	Input []string `yaml:"input"`
}

func (a *Action) Exec(input []string, state State) {
}
