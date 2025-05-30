package models

import (
	"os"
	"syscall"

	"github.com/Rokkit-exe/rokkitland/tui"
	"golang.org/x/term"
)

type InputManager struct{}

func (i *InputManager) RecordKeys(state *State) error {
	buf := make([]byte, 3)
	os.Stdin.Read(buf)

	if buf[0] == tui.Escape1 && buf[1] == tui.Escape2 { // Arrow keys
		switch buf[2] {
		case tui.Up: // Up
			state.MoveCursorUp()
		case tui.Down: // Down
			state.MoveCursorDown()
		case tui.Left: // Left
			state.MoveCursorLeft()
		case tui.Right: // Right
			state.MoveCursorRight()
		}
	}
	switch buf[0] {
	case tui.Enter:
		state.SelectSection()
	case tui.Space:
		state.ToggleSelectOption()
	case tui.Quit:
		i.Quit(state)
	case tui.Tab:
		state.TogglePage()
	case tui.Install:
		state.InstallSelectedOptions()
	case tui.Remove:
		state.RemoveSelectedOptions()
	case tui.Toggle:
		state.ToggleAllOptions()
	default:
	}

	return nil
}

func (i *InputManager) Quit(state *State) {
	state.Console.Add("--------------------------------------------------------")
	state.Console.Add("Exiting...")
	term.Restore(int(syscall.Stdin), state.OldState)
	os.Exit(0)
}
