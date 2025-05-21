package models

type Action struct {
	Name        string
	Description string
	Cmd         func([]string) error
	Input       []string
}

func (a *Action) Exec() error {
	return a.Cmd(a.Input)
}
