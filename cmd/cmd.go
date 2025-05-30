package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Rokkit-exe/rokkitland/models"
)

var funcmap = map[string]func([]string, models.State){
	"install-packages":   InstallPackages,
	"uninstall_packages": UninstallPackages,
	"exec_scripts":       ExecScripts,
}

func Exec(action *models.Action, input []string, state models.State) error {
	var err error

	if f, ok := funcmap[action.Cmd]; ok {
		f(input, state)
	} else {
		err = fmt.Errorf("Error: Unknown command: %s", action.Cmd)
		state.Log.Add(err.Error())
	}
	return err
}

func InstallPackage(packageName string) error {
	command := exec.Command("yay", "-S", "--noconfirm", "--needed", packageName)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return fmt.Errorf("Failed to install %s: %w", packageName, err)
	}
	return nil
}

func UninstallPackage(packageName string) error {
	command := exec.Command("yay", "-Rns", "--noconfirm", "--needed", packageName)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return fmt.Errorf("Failed to uninstall %s: %w", packageName, err)
	}
	return nil
}

func InstallPackages(packages []string, state models.State) {
	for i, pkg := range packages {
		msg := fmt.Sprintf("Installing package %s/%s", i+1, len(packages))
		state.Log.Add(msg)
		err := InstallPackage(pkg)
		if err != nil {
			state.Log.Add(err.Error())
		} else {
			success := fmt.Sprintf("Successfully installed: %s", pkg)
			state.Log.Add(success)
		}
	}
}

func UninstallPackages(packages []string, state models.State) {
	for i, pkg := range packages {
		msg := fmt.Sprintf("Uninstalling package %s/%s", i+1, len(packages))
		state.Log.Add(msg)
		err := UninstallPackage(pkg)
		if err != nil {
			state.Log.Add(err.Error())
		} else {
			success := fmt.Sprintf("Successfully uninstalled: %s", pkg)
			state.Log.Add(success)
		}
	}
}

func ExecScript(script string) error {
	command := exec.Command("bash", script)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return fmt.Errorf("failed to execute script %s: %w", script, err)
	}
	return nil
}

func ExecScripts(scripts []string, state models.State) {
	for i, script := range scripts {
		msg := fmt.Sprintf("Executing script %s/%s", i+1, len(scripts))
		state.Log.Add(msg)
		err := ExecScript(script)
		if err != nil {
			state.Log.Add(err.Error())
		} else {
			success := fmt.Sprintf("Successfully executed: %s", script)
			state.Log.Add(success)
		}
	}
}
