package controller

import (
	"os/exec"
	"strings"

	"github.com/Rokkit-exe/rokkitland/models"
)

type ConsoleController struct {
	State *models.State
}

func NewConsoleController(state *models.State) *ConsoleController {
	return &ConsoleController{
		State: state,
	}
}

func (c *ConsoleController) RunCommandWithPTY(cmdArgs []string) {
	c.State.Console.Add("[info] running: " + strings.Join(cmdArgs, " "))

	ouput, err := exec.Command(cmdArgs[0], cmdArgs[1:]...).CombinedOutput()

	if err != nil {
		c.State.Console.Add("[error] " + err.Error())
	} else {
		c.State.Console.Add("[info] command output: " + string(ouput))
	}

	c.State.Console.Add(string(ouput))
}
