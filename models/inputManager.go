package models

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/Rokkit-exe/rokkitland/tui"
	"golang.org/x/term"
)

type InputManager struct {
	State State
}

func (i *InputManager) RecordKeys() error {
	buf := make([]byte, 3)
	os.Stdin.Read(buf)

	if buf[0] == tui.Enter { // Enter
		i.State.SelectSection()
	}
	if buf[0] == tui.Escape1 && buf[1] == tui.Escape2 { // Arrow keys
		switch buf[2] {
		case tui.Up: // Up
			i.State.MoveCursorUp()
		case tui.Down: // Down
			i.State.MoveCursorDown()
		case tui.Left: // Left
			i.State.MoveCursorLeft()
		case tui.Right: // Right
			i.State.MoveCursorRight()
		}
	}
	if buf[0] == tui.Space {
		i.State.ToggleSelectOption()
	}

	if buf[0] == tui.Quit { // Escape1
		i.Quit()
	}
	if buf[0] == tui.Tab { // Tab
		i.State.ToggleSelectedPanel()
	}
	return nil
}

func (i *InputManager) Quit() {
	i.State.MoveCursor(90, 0)
	fmt.Println("--------------------------------------------------------")
	fmt.Println("Exiting...")
	time.Sleep(1 * time.Second)
	i.State.Clear()
	term.Restore(int(syscall.Stdin), i.State.OldState)
	os.Exit(0)
}
