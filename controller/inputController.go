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

func (i *InputController) RecordInput() error {
	buf := make([]byte, 100)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		i.State.Console.Add("[error] Failed to read input: "+err.Error(), tui.Red)
		return err
	}
	if i.State.IsCommandRunning {
		i.RecordCommandInput(buf[:n])
	} else {
		i.RecordUiInput(buf[:n])
	}
	return nil
}

func (i *InputController) RecordUiInput(buf []byte) {
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
		case tui.CtrlC:
			i.ConsoleController.RestoreUIMode()
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
}

func (i *InputController) RecordCommandInput(buf []byte) {
	i.State.CommandInputChan <- buf
}

func (i *InputController) Quit() {
	i.State.Console.Add("--------------------------------------------------------", tui.Green)
	i.State.Console.Add("Exiting...", tui.Green)
	term.Restore(int(syscall.Stdin), i.State.OldState)
	os.Exit(0)
}
