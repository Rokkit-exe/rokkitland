package models

import (
	"fmt"
	"strings"

	"github.com/Rokkit-exe/rokkitland/tui"
)

type Page struct {
	Id       int       `yaml:"id"`
	Title    string    `yaml:"title"`
	Panels   []Panel   `yaml:"panels"`
	Sections []Section `yaml:"sections,omitempty"`
}

func (p *Page) DrawTab(state *State, active bool) {
	posX := state.Cursor.X
	if active {
		fmt.Printf("%s", tui.Green.ANSI())
	} else {
		fmt.Printf("%s", tui.Reset.ANSI())
	}
	fmt.Printf("┌%s┐", strings.Repeat("─", len(p.Title)+2))
	state.MoveCursor(state.Cursor.Y+1, posX)
	fmt.Printf("%s %s %s", "│", p.Title, "│")
	state.MoveCursor(state.Cursor.Y+1, posX)
	fmt.Printf("└%s┘", strings.Repeat("─", len(p.Title)+2))
	fmt.Printf("%s", tui.Reset.ANSI())
	state.MoveCursor(1, posX+len(p.Title)+4)
}

func (p *Page) DrawPanels(state *State) {
	for _, panel := range p.Panels {
		panel.Draw(state)
	}
}
