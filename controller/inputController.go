package controller

import (
	"os"
	"syscall"

	"github.com/Rokkit-exe/rokkitland/models"
	"github.com/Rokkit-exe/rokkitland/tui"
	"golang.org/x/term"
)

type InputController struct {
	State             *models.State
	StateController   *StateController
	ConsoleController *ConsoleController
}

func NewInputController(state *models.State, stateController *StateController, consoleController *ConsoleController) *InputController {
	return &InputController{
		State:             state,
		StateController:   stateController,
		ConsoleController: consoleController,
	}
}

func (i *InputController) RecordKeys() error {
	buf := make([]byte, 3)
	_, err := os.Stdin.Read(buf)
	if err != nil {
		return err
	}

	// if i.State.IsCommandRunning {
	// 	// Forward to command PTY
	// 	for _, b := range buf {
	// 		if b != 0 {
	// 			i.State.CommandInputChan <- b
	// 		}
	// 	}
	// 	return nil
	// }

	// Normal TUI controls
	if buf[0] == tui.Escape1 && buf[1] == tui.Escape2 {
		switch buf[2] {
		case tui.Up:
			i.StateController.MoveCursorUp()
		case tui.Down:
			i.StateController.MoveCursorDown()
		case tui.Left:
			i.StateController.MoveCursorLeft()
		case tui.Right:
			i.StateController.MoveCursorRight()
		}
	} else {
		switch buf[0] {
		case tui.Enter:
			i.StateController.SelectSection()
		case tui.Space:
			i.StateController.ToggleSelectOption()
		case tui.Quit:
			i.Quit()
		case tui.Tab:
			i.StateController.TogglePage()
		case tui.Install:
			i.StateController.InstallSelectedOptions()
		case tui.Remove:
			i.StateController.RemoveSelectedOptions()
		case tui.Toggle:
			i.StateController.ToggleAllOptions()
		}
	}

	return nil
}

func (i *InputController) Quit() {
	i.State.Console.Add("--------------------------------------------------------")
	i.State.Console.Add("Exiting...")
	term.Restore(int(syscall.Stdin), i.State.OldState)
	os.Exit(0)
}
