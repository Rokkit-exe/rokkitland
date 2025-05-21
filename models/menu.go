package models

import (
	"syscall"

	"golang.org/x/term"
)

type Menu struct {
	Title        string
	Pages        []Page
	InputManager InputManager
	Log          Log
	State        State
}

func (m *Menu) DrawMenu() error {
	oldstate, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		return err
	}
	m.State.OldState = oldstate
	defer term.Restore(int(syscall.Stdin), m.State.OldState)

	for {
		m.State.Clear()
		m.State.MoveCursor(0, 0)
		for _, page := range m.State.Pages {
			page.DrawTab(m.State, true)
		}
		m.State.MoveCursor(3, 0)
		m.State.Pages[m.State.SelectedPage].DrawPanels(m.InputManager)
	}
}
