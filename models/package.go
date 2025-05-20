package models

import (
	"fmt"
	"os/exec"
	"strings"
)

type Package struct {
	Name        string `json:"name"`
	Selected    bool   `json:"selected"`
	Description string `json:"description"`
}

func (p Package) String() string {
	return fmt.Sprintf("%s | %s", p.Name, p.Description)
}

func (p *Package) GetDescription() {
	cmd := exec.Command("pacman", "-Qi", p.Name, "|", "grep", "Description")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	description := strings.TrimPrefix(string(output), "Description     : ")
	p.Description = description
}
