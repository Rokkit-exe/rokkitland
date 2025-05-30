package models

import (
	"fmt"
	"html"
	"os/exec"

	"github.com/Rokkit-exe/rokkitland/utils"
)

type Option struct {
	Name        string `json:"name"`
	Selected    bool   `json:"selected"`
	Description string `json:"description"`
}

func (o Option) String() string {
	return fmt.Sprintf("%s | %s", o.Name, o.Description)
}

func (o *Option) GetDescription() string {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("yay -Si %s | grep Description", o.Name))
	output, err := cmd.Output()
	fmt.Println("Command:", cmd.String())
	if err != nil {
		fmt.Println("Error:", err)
		return "no description available"
	}
	description := utils.TrimUntil(string(output), ':')
	description = html.UnescapeString(description)
	return description
}
