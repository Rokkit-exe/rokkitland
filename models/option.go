package models

import (
	"fmt"
	"os/exec"
	"strings"
)

type Option struct {
	Name        string `json:"name"`
	Selected    bool   `json:"selected"`
	Description string `json:"description"`
}

func (o Option) String() string {
	return fmt.Sprintf("%s | %s", o.Name, o.Description)
}

func (o *Option) GetDescription() {
	cmd := exec.Command("pacman", "-Qi", o.Name, "|", "grep", "Description")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	description := strings.TrimPrefix(string(output), "Description     : ")
	o.Description = description
}
