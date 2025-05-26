package models

import (
	"fmt"
	"strings"

	"github.com/Rokkit-exe/rokkitland/tui"
)

type Panel struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Format   string `json:"format"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	PaddingX int    `json:"padding-x"`
	PaddingY int    `json:"padding-y"`
}

func (p *Panel) DrawBox(state *State, active bool) {
	if active {
		fmt.Printf("%s", tui.Green.ANSI())
	} else {
		fmt.Printf("%s", tui.Reset.ANSI())
	}
	state.MoveCursor(p.Y, p.X)
	fmt.Printf("┌─%s%s┐", p.Title, strings.Repeat("─", p.Width-len(p.Title)-2))
	for i := 1; i < p.Height-1; i++ {
		state.MoveCursor(p.Y+i, p.X)
		fmt.Print("│" + strings.Repeat(" ", p.Width-1) + "│")
	}
	state.MoveCursor(p.Y+p.Height-1, p.X)
	fmt.Printf("└%s┘", strings.Repeat("─", p.Width-1))
	fmt.Printf("%s", tui.Reset.ANSI())
}

func (p *Panel) DrawActionPanel(state *State) {
	if state.SelectedPanel == 3 {
		p.DrawBox(state, true)
	} else {
		p.DrawBox(state, false)
	}
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)
	// right Panel
	for i, act := range state.Actions {
		cursorPrefix := "  "
		if i == state.ActionCursor && state.SelectedPanel == 3 {
			state.ActionCursor = i
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		state.MoveCursor(p.Y+p.PaddingY+i, p.X+p.PaddingX)
		fmt.Printf("%s%s", cursorPrefix, act.Name)
	}
}

func (p *Panel) DrawNavPanel(state *State) {
	p.DrawBox(state, false)
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)
	fmt.Printf("Press ↑ ↓ to navigate between sections and options")
	state.MoveCursor(p.Y+p.PaddingY+1, p.X+p.PaddingX)
	fmt.Printf("Press ␣ to toggle an option")
	state.MoveCursor(p.Y+p.PaddingY+2, p.X+p.PaddingX)
	fmt.Printf("Press ↵ to select a section")
	state.MoveCursor(p.Y+p.PaddingY+3, p.X+p.PaddingX)
	fmt.Printf("Press ⇥ to switch between panels")
}

func (p *Panel) DrawOptionPanel(state *State) {
	if state.SelectedPanel == 2 {
		p.DrawBox(state, true)
	} else {
		p.DrawBox(state, false)
	}
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)
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
		state.MoveCursor(p.Y+p.PaddingY+i, p.X+p.PaddingX)
		fmt.Printf("%s%s %s", cursorPrefix, prefix, opt.Name)
	}
}

func (p *Panel) DrawSectionPanel(state *State) {
	if state.SelectedPanel == 1 {
		p.DrawBox(state, true)
	} else {
		p.DrawBox(state, false)
	}
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)
	// left Panels[1]
	for i, opt := range state.Sections {
		cursorPrefix := "  "
		if i == state.SectionCursor && state.SelectedPanel == 1 {
			state.SelectedSection = i
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		state.MoveCursor(p.Y+p.PaddingY+i, p.X+p.PaddingX)
		fmt.Printf("%s%s", cursorPrefix, opt.Title)
	}
}

func (p *Panel) DrawLogPanel(state *State) {
	p.DrawBox(state, false)
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)

	n := 5
	lenLog := len(state.Log)

	if lenLog == 0 {
		return
	} else if lenLog > 0 && lenLog < p.Height-2 {
		n = lenLog
	}

	for i, log := range state.Log.LastN(n) {
		fmt.Printf("%s", log)
		state.MoveCursor(p.Y+p.PaddingY+i, p.X+p.PaddingX)
	}
}

func (p *Panel) Draw(state *State) {
	switch p.Format {
	case "nav":
		p.DrawNavPanel(state)
	case "option":
		p.DrawOptionPanel(state)
	case "action":
		p.DrawActionPanel(state)
	case "section":
		p.DrawSectionPanel(state)
	case "log":
		p.DrawLogPanel(state)
	default:
		p.DrawBox(state, false)
	}
}
