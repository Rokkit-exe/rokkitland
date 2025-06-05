package controller

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Rokkit-exe/rokkitland/models"
	"github.com/Rokkit-exe/rokkitland/tui"
	"github.com/creack/pty"
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
	// c.State.CreateCommandInputChan()
	cmd := exec.Command("script", "-qfc", strings.Join(cmdArgs, " "), "/dev/null")
	ptmx, err := pty.Start(cmd)
	if err != nil {
		c.State.Console.Add("[error] PTY start failed: " + err.Error())
		return
	}
	defer func(ptmx *os.File) {
		c.State.Console.Add("[info] Command Executed")
		c.RestoreUIMode(ptmx)
	}(ptmx)

	// Read output asynchronously
	go c.ReadFromPTY(ptmx)

	go c.WriteToPTY(ptmx)

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		c.State.Console.Add("[error] command finished with error: " + err.Error())
	}
}

func (c *ConsoleController) WriteToPTY(ptmx *os.File) {
	if c.State.CommandInputChan == nil {
		c.State.Console.Add("[error] CommandInputChan is nil, cannot read input")
		return
	}
	for chunk := range c.State.CommandInputChan {
		for _, b := range chunk {
			_, err := ptmx.Write([]byte{b})
			if err != nil {
				c.State.Console.Add("[error] failed to write to PTY: " + err.Error())
				return
			}
			if b == tui.CtrlC {
				c.State.Console.Add("[input] Ctrl+C detected, restoring UI mode")
				c.RestoreUIMode(ptmx)
				return
			}
		}
	}
}

func (c *ConsoleController) ReadFromPTY(ptmx *os.File) {
	c.State.Console.Add("[info] Reading from PTY...")
	buf := make([]byte, 1024)
	for {
		n, err := ptmx.Read(buf)
		if n > 0 {
			c.State.Console.Add(string(buf[:n]))
		}
		if err != nil {
			c.State.Console.Add("[error] failed to read from PTY: " + err.Error())
			break
		}
	}
}

func (c *ConsoleController) RestoreUIMode(ptmx *os.File) {
	_ = ptmx.Close() // Close the PTY file descriptor
	c.State.SetIsCommandRunning(false)
	c.State.SelectedPanel = 1
	close(c.State.CommandInputChan) // Close the command input channel
	c.State.CommandInputChan = nil  // Clear the command input channel
	c.State.Console.Add("[info] Restored UI mode")
}
