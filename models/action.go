package models

import (
	"fmt"

	"github.com/Rokkit-exe/rokkitland/cmd"
)

type Action struct {
	Title string   `yaml:"title"`
	Cmd   string   `yaml:"cmd"`
	Input []string `yaml:"input"`
}

func (a *Action) Exec(input []string, state State) error {
	funcmap := map[string]func([]string, State){
		"install-packages":   cmd.InstallPackages,
		"uninstall_packages": cmd.UninstallPackages,
		"exec_scripts":       cmd.ExecScripts,
	}

	var err error

	if f, ok := funcmap[a.Cmd]; ok {
		f(input, state)
	} else {
		err = fmt.Errorf("Error: Unknown command: %s", a.Cmd)
		state.Log.Add(err.Error())
	}

	return err
}
