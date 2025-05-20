package install

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallPackage(packageName string) error {
	cmd := exec.Command("yay", "-S", "--noconfirm", "--needed", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install %s: %w", packageName, err)
	}
	return nil
}

func InstallPackages(packages []string) {
	for _, pkg := range packages {
		err := InstallPackage(pkg)
		if err != nil {
			fmt.Printf("Error installing %s: %v\n", pkg, err)
		} else {
			fmt.Printf("Successfully installed %s\n", pkg)
		}
	}
}
