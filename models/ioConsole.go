package models

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

type IOConsole struct {
	Lines      []string
	WriteMutex sync.Mutex
}

func (c *IOConsole) Add(line string) {
	c.WriteMutex.Lock()
	defer c.WriteMutex.Unlock()
	c.Lines = append(c.Lines, line)
	if len(c.Lines) > 200 {
		c.Lines = c.Lines[len(c.Lines)-200:]
	}
}

func (c *IOConsole) GetLines() []string {
	c.WriteMutex.Lock()
	defer c.WriteMutex.Unlock()
	return append([]string(nil), c.Lines...)
}

func (c *IOConsole) LastN(n int) []string {
	if len(c.Lines) < n {
		return c.Lines
	}
	return c.Lines[len(c.Lines)-n:]
}

func (c *IOConsole) RunCommandInPanel(cmdArgs []string, inputChan <-chan byte, doneChan chan<- struct{}) {
	if len(cmdArgs) == 0 {
		c.Add("[error] no command provided")
		close(doneChan)
		return
	}

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		c.Add("[error] failed to attach stdin")
		close(doneChan)
		return
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		c.Add("[error] failed to attach stdout")
		close(doneChan)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		c.Add("[error] failed to attach stderr")
		close(doneChan)
		return
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		c.Add(fmt.Sprintf("[error] failed to start: %s", err))
		close(doneChan)
		return
	}

	// Pipe output into console
	go func() {
		scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
		for scanner.Scan() {
			c.Add(scanner.Text())
		}
	}()

	// Forward keyboard input into command
	go func() {
		for b := range inputChan {
			stdin.Write([]byte{b})
		}
	}()

	// Wait for command to finish
	go func() {
		_ = cmd.Wait()
		close(doneChan)
	}()
}
