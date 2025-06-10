package controller

import (
	"errors"
	"io"
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
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		c.State.Console.Add("[error] PTY start failed: "+err.Error(), tui.Red)
		return
	}
	c.State.Ptmx = ptmx
	defer func() {
		c.RestoreUIMode()
	}()

	go c.Read()

	go c.Write()

	err = cmd.Wait()
	if err != nil {
		c.State.Console.Add("[error] command finished with error: "+err.Error(), tui.Red)
	}
}

func (c *ConsoleController) Write() {
	if c.State.CommandInputChan == nil {
		c.State.Console.Add("[error] CommandInputChan is nil, cannot read input", tui.Red)
		return
	}
	for chunk := range c.State.CommandInputChan {
		for _, b := range chunk {
			_, err := c.State.Ptmx.Write([]byte{b})
			if err != nil {
				c.State.Console.Add("[error] failed to write to PTY: "+err.Error(), tui.Red)
				return
			}
			if b == tui.CtrlC {
				c.State.Console.Add("[input] Ctrl+C detected, restoring UI mode", tui.Blue)
				c.RestoreUIMode()
				return
			}
		}
	}
}

func (c *ConsoleController) Read() {
	buf := make([]byte, 1024)
	for {
		n, err := c.State.Ptmx.Read(buf)
		if n > 0 {
			c.State.Console.Add(string(buf[:n]), tui.White)
		}
		if err != nil {
			if errors.Is(err, io.EOF) || strings.Contains(err.Error(), "input/output error") {
				c.State.Console.Add("......................................................................", tui.Blue)
			} else {
				c.State.Console.Add("[error] failed to read from PTY: "+err.Error(), tui.Red)
			}
			break
		}
	}
}

func (c *ConsoleController) SetCommandMode() {
	c.State.SetIsCommandRunning(true)
	c.State.SelectedPanel = 5
	c.State.CommandInputChan = make(chan []byte)
}

func (c *ConsoleController) RestoreUIMode() {
	_ = c.State.Ptmx.Close() // Close the PTY file descriptor
	c.State.Ptmx = nil       // Clear the PTY file descriptor
	c.State.SetIsCommandRunning(false)
	c.State.SelectedPanel = 1
	close(c.State.CommandInputChan) // Close the command input channel
	c.State.CommandInputChan = nil  // Clear the command input channel
	c.State.Console.Add("[info] UI Mode Restored", tui.Blue)
}
