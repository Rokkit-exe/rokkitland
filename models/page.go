package models

import (
	"fmt"
	"github.com/Rokkit-exe/rokkitland/tui"
	"strings"
)

type Page struct {
	Title  string  `json:"title"`
	Panels []Panel `json:"panels"`
	State  State   `json:"state"`
}

func (p *Page) DrawTab(state State, active bool) {
	if active {
		fmt.Printf("%s", tui.Green.ANSI())
	} else {
		fmt.Printf("%s", tui.Reset.ANSI())
	}
	fmt.Printf("┌%s┐", strings.Repeat("─", len(p.Title)+2))
	state.MoveCursor(p.State.Cursor.Y+1, p.State.Cursor.X)
	fmt.Print("│" + p.Title + "│")
	state.MoveCursor(p.State.Cursor.Y+1, p.State.Cursor.X)
	fmt.Printf("└%s┘", strings.Repeat("─", len(p.Title)+2))
	fmt.Printf("%s", tui.Reset.ANSI())
	state.MoveCursor(0, state.Cursor.X)
}

func (p *Page) DrawPanels(inputManager InputManager) error {
	for {
		p.State.Clear()
		for _, panel := range p.Panels {
			panel.Draw(p.State)
		}

		err := inputManager.RecordKeys()
		if err != nil {
			p.State.MoveCursor(90, 0)
			fmt.Println("Error:", err)
			return err
		}
	}
}
