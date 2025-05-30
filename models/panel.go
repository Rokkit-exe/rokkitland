package models

import (
	"fmt"
	"strings"

	"github.com/Rokkit-exe/rokkitland/tui"
	"github.com/Rokkit-exe/rokkitland/utils"
)

type Panel struct {
	Id       int    `yaml:"id"`
	Title    string `yaml:"title"`
	Format   string `yaml:"format"`
	X        int    `yaml:"x"`
	Y        int    `yaml:"y"`
	Width    int    `yaml:"width"`
	Height   int    `yaml:"height"`
	PaddingX int    `yaml:"padding-x"`
	PaddingY int    `yaml:"padding-y"`
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
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)
	p.DrawBox(state, false)
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)
	fmt.Printf("Press I to install selected options")
	state.MoveCursor(p.Y+p.PaddingY+1, p.X+p.PaddingX)
	fmt.Printf("Press R to remove selected options")
	state.MoveCursor(p.Y+p.PaddingY+2, p.X+p.PaddingX)
	fmt.Printf("Press H for help")
	state.MoveCursor(p.Y+p.PaddingY+3, p.X+p.PaddingX)
	fmt.Printf("Press Q to quit")
}

func (p *Panel) DrawNavPanel(state *State) {
	p.DrawBox(state, false)
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)
	fmt.Printf("Press ↑ ↓ to move cursor")
	state.MoveCursor(p.Y+p.PaddingY+1, p.X+p.PaddingX)
	fmt.Printf("Press ␣ to toggle an option")
	state.MoveCursor(p.Y+p.PaddingY+2, p.X+p.PaddingX)
	fmt.Printf("Press ↵ to select a section")
	state.MoveCursor(p.Y+p.PaddingY+3, p.X+p.PaddingX)
	fmt.Printf("Press ⇥ to toggle pages")
}

func (p *Panel) DrawOptionPanel(state *State) {
	if state.SelectedPanel == 2 {
		p.DrawBox(state, true)
	} else {
		p.DrawBox(state, false)
	}
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)

	options := &state.Pages[state.SelectedPage].Sections[state.SelectedSection].Options
	if len(*options) == 0 {
		state.Console.Add("No options available in this section.")
		return
	}
	for i, opt := range *options {
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

func (p *Panel) DrawDescriptionPanel(state *State) {
	p.DrawBox(state, false)
	option := &state.Pages[state.SelectedPage].Sections[state.SelectedSection].Options[state.OptionCursor]
	lines := utils.WrapWords(option.Description, p.Width-p.PaddingX*2)
	for i, line := range lines {
		state.MoveCursor(p.Y+p.PaddingY+i, p.X+p.PaddingX)
		fmt.Printf("%s", line)
	}
}

func (p *Panel) DrawSectionPanel(state *State) {
	if state.SelectedPanel == 1 {
		p.DrawBox(state, true)
	} else {
		p.DrawBox(state, false)
	}
	state.MoveCursor(p.Y+p.PaddingY, p.X+p.PaddingX)
	sections := &state.Pages[state.SelectedPage].Sections
	for i, opt := range *sections {
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

	n := p.Height - 2 // Number of logs to display
	lenLog := len(state.Console.Lines)

	if lenLog == 0 {
		return
	} else if lenLog > 0 && lenLog < p.Height-2 {
		n = lenLog
	}

	for i, log := range state.Console.LastN(n) {
		fmt.Printf("%s", log)
		state.MoveCursor(p.Y+p.PaddingY+i+1, p.X+p.PaddingX)
	}
}

func (p *Panel) Draw(state *State) {
	switch p.Format {
	case "nav":
		p.DrawNavPanel(state)
	case "option":
		p.DrawOptionPanel(state)
	case "description":
		p.DrawDescriptionPanel(state)
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
