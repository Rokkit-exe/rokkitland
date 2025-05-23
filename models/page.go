package models

import (
	"fmt"
	"strings"

	"github.com/Rokkit-exe/rokkitland/tui"
)

type Page struct {
	Title  string  `json:"title"`
	Panels []Panel `json:"panels"`
}

func (p *Page) DrawTab(state *State, active bool) {
	if active {
		fmt.Printf("%s", tui.Green.ANSI())
	} else {
		fmt.Printf("%s", tui.Reset.ANSI())
	}
	fmt.Printf("┌%s┐", strings.Repeat("─", len(p.Title)+2))
	state.MoveCursor(state.Cursor.Y+1, 0)
	fmt.Print("│" + p.Title + "│")
	state.MoveCursor(state.Cursor.Y+1, 0)
	fmt.Printf("└%s┘", strings.Repeat("─", len(p.Title)+2))
	fmt.Printf("%s", tui.Reset.ANSI())
	state.MoveCursor(0, len(p.Title)+2)
}

func (p *Page) DrawPanels(state *State, inputManager InputManager) error {
	for {
		state.Clear()
		for _, panel := range p.Panels {
			panel.Draw(state)
		}

		err := inputManager.RecordKeys()
		if err != nil {
			state.MoveCursor(90, 0)
			fmt.Println("Error:", err)
			return err
		}
	}
}
