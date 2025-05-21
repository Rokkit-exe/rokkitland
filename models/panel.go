package models

import (
	"fmt"
	"strings"

	"github.com/Rokkit-exe/rokkitland/tui"
)

type position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Panel struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	Format     string     `json:"format"`
	Content    []string   `json:"content"`
	Pos        position   `json:"pos"`
	Dimensions dimensions `json:"dimensions"`
}

func (p *Panel) DrawBox(state State, active bool) {
	if active {
		fmt.Printf("%s", tui.Green.ANSI())
	} else {
		fmt.Printf("%s", tui.Reset.ANSI())
	}
	state.MoveCursor(p.Pos.Y, p.Pos.X)
	fmt.Printf("┌─%s%s┐", p.Title, strings.Repeat("─", p.Dimensions.Width-len(p.Title)-2))
	for i := 1; i < p.Dimensions.Height-1; i++ {
		state.MoveCursor(p.Pos.Y+i, p.Pos.X)
		fmt.Print("│" + strings.Repeat(" ", p.Dimensions.Width-1) + "│")
	}
	state.MoveCursor(p.Pos.Y+p.Dimensions.Height-1, p.Pos.X)
	fmt.Printf("└%s┘", strings.Repeat("─", p.Dimensions.Width-1))
	fmt.Printf("%s", tui.Reset.ANSI())
}

func (p *Panel) DrawNavPanel(state State) {
	p.DrawBox(state, false)
	state.MoveCursor(p.Pos.Y+1, p.Pos.X+3)
	for i, msg := range p.Content {
		state.MoveCursor(p.Pos.Y+1+i, p.Pos.X+3)
		fmt.Print(msg)
	}
}

func (p *Panel) DrawSectionPanel(state State) {
	if state.SelectedPanel == 1 {
		p.DrawBox(state, true)
	} else {
		p.DrawBox(state, false)
	}
	state.MoveCursor(p.Pos.Y+1, p.Pos.X+3)
	// left Panels[1]
	for i, opt := range state.Sections {
		cursorPrefix := "  "
		if i == state.SectionCursor && state.SelectedPanel == 1 {
			state.SelectedSection = i
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		state.MoveCursor(p.Pos.Y+1+i, p.Pos.X+3)
		fmt.Printf("%s%s", cursorPrefix, opt.Title)
	}
}

func (p *Panel) DrawActionsPanel(state State) {
	if state.SelectedPanel == 3 {
		p.DrawBox(state, true)
	} else {
		p.DrawBox(state, false)
	}
	state.MoveCursor(p.Pos.Y+1, p.Pos.X+3)
	for i, opt := range p.Content {
		cursorPrefix := "  "
		if i == state.SectionCursor && state.SelectedPanel == 3 {
			state.SelectedSection = i
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		state.MoveCursor(p.Pos.Y+1+i, p.Pos.X+3)
		fmt.Printf("%s%s", cursorPrefix, opt)
	}
}

func (p *Panel) DrawOptionsPanel(state State) {
	if state.SelectedPanel == 2 {
		p.DrawBox(state, true)
	} else {
		p.DrawBox(state, false)
	}
	state.MoveCursor(p.Pos.Y+1, p.Pos.X+3)
	// right Panel
	for i, opt := range state.Sections[state.SelectedSection].Options {
		prefix := "[ ]"
		if opt.Selected {
			prefix = fmt.Sprintf("[%sx%s]", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		cursorPrefix := "  "
		if i == state.OptionCursor && state.SelectedPanel == 2 {
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		state.MoveCursor(p.Pos.Y+1+i, p.Pos.X+3)
		fmt.Printf("%s%s %s", cursorPrefix, prefix, opt.Name)
	}
}

func (p *Panel) DrawLogPanel(state State) {
	p.DrawBox(state, false)
	state.MoveCursor(p.Pos.Y+1, p.Pos.X+3)
}

func (p *Panel) Draw(state State) {
	switch p.Format {
	case "nav":
		p.DrawNavPanel(state)
	case "section":
		p.DrawSectionPanel(state)
	case "option":
		p.DrawOptionsPanel(state)
	case "action":
		p.DrawActionsPanel(state)
	case "log":
		p.DrawLogPanel(state)
	}
}
