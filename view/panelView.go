package view

import (
	"fmt"
	"strings"

	"github.com/Rokkit-exe/rokkitland/models"
	"github.com/Rokkit-exe/rokkitland/tui"
	"github.com/Rokkit-exe/rokkitland/utils"
)

type PanelView struct {
	State *models.State
}

func NewPanelView(state *models.State) *PanelView {
	return &PanelView{
		State: state,
	}
}

func (p *PanelView) DrawBox(panel *models.Panel, active bool) {
	if active {
		fmt.Printf("%s", tui.Green.ANSI())
	} else {
		fmt.Printf("%s", tui.Reset.ANSI())
	}
	p.State.Cursor.Move(panel.Y, panel.X)
	fmt.Printf("┌─%s%s┐", panel.Title, strings.Repeat("─", panel.Width-len(panel.Title)-2))
	for i := 1; i < panel.Height-1; i++ {
		p.State.Cursor.Move(panel.Y+i, panel.X)
		fmt.Print("│" + strings.Repeat(" ", panel.Width-1) + "│")
	}
	p.State.Cursor.Move(panel.Y+panel.Height-1, panel.X)
	fmt.Printf("└%s┘", strings.Repeat("─", panel.Width-1))
	fmt.Printf("%s", tui.Reset.ANSI())
}

func (p *PanelView) DrawActionPanel(panel *models.Panel) {
	p.State.Cursor.Move(panel.Y+panel.PaddingY, panel.X+panel.PaddingX)
	p.DrawBox(panel, false)
	p.State.Cursor.Move(panel.Y+panel.PaddingY, panel.X+panel.PaddingX)
	fmt.Printf("Press I to install selected options")
	p.State.Cursor.Move(panel.Y+panel.PaddingY+1, panel.X+panel.PaddingX)
	fmt.Printf("Press R to remove selected options")
	p.State.Cursor.Move(panel.Y+panel.PaddingY+2, panel.X+panel.PaddingX)
	fmt.Printf("Press H for help")
	p.State.Cursor.Move(panel.Y+panel.PaddingY+3, panel.X+panel.PaddingX)
	fmt.Printf("Press Q to quit")
}

func (p *PanelView) DrawNavPanel(panel *models.Panel) {
	p.DrawBox(panel, false)
	p.State.Cursor.Move(panel.Y+panel.PaddingY, panel.X+panel.PaddingX)
	fmt.Printf("Press ↑ ↓ to move cursor")
	p.State.Cursor.Move(panel.Y+panel.PaddingY+1, panel.X+panel.PaddingX)
	fmt.Printf("Press ␣ to toggle an option")
	p.State.Cursor.Move(panel.Y+panel.PaddingY+2, panel.X+panel.PaddingX)
	fmt.Printf("Press ↵ to select a section")
	p.State.Cursor.Move(panel.Y+panel.PaddingY+3, panel.X+panel.PaddingX)
	fmt.Printf("Press ⇥ to toggle pages")
}

func (p *PanelView) DrawOptionPanel(panel *models.Panel) {
	if p.State.SelectedPanel == 2 {
		p.DrawBox(panel, true)
	} else {
		p.DrawBox(panel, false)
	}
	p.State.Cursor.Move(panel.Y+panel.PaddingY, panel.X+panel.PaddingX)

	options := &p.State.Pages[p.State.SelectedPage].Sections[p.State.SelectedSection].Options
	for i, opt := range *options {
		prefix := "[ ]"
		if opt.Selected {
			prefix = fmt.Sprintf("[%sx%s]", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		cursorPrefix := "  "
		if i == p.State.OptionCursor && p.State.SelectedPanel == 2 {
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		p.State.Cursor.Move(panel.Y+panel.PaddingY+i, panel.X+panel.PaddingX)
		fmt.Printf("%s%s %s", cursorPrefix, prefix, opt.Name)
	}
}

func (p *PanelView) DrawDescriptionPanel(panel *models.Panel) {
	p.DrawBox(panel, false)
	option := &p.State.Pages[p.State.SelectedPage].Sections[p.State.SelectedSection].Options[p.State.OptionCursor]
	lines := utils.WrapWords(option.Description, panel.Width-panel.PaddingX*2)
	for i, line := range lines {
		p.State.Cursor.Move(panel.Y+panel.PaddingY+i, panel.X+panel.PaddingX)
		fmt.Printf("%s", line)
	}
}

func (p *PanelView) DrawSectionPanel(panel *models.Panel) {
	if p.State.SelectedPanel == 1 {
		p.DrawBox(panel, true)
	} else {
		p.DrawBox(panel, false)
	}
	p.State.Cursor.Move(panel.Y+panel.PaddingY, panel.X+panel.PaddingX)
	sections := &p.State.Pages[p.State.SelectedPage].Sections
	for i, opt := range *sections {
		cursorPrefix := "  "
		if i == p.State.SectionCursor && p.State.SelectedPanel == 1 {
			p.State.SelectedSection = i
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		p.State.Cursor.Move(panel.Y+panel.PaddingY+i, panel.X+panel.PaddingX)
		fmt.Printf("%s%s", cursorPrefix, opt.Title)
	}
}

func (p *PanelView) DrawLogPanel(panel *models.Panel) {
	if p.State.SelectedPanel == 5 {
		p.DrawBox(panel, true)
	} else {
		p.DrawBox(panel, false)
	}
	p.State.Cursor.Move(panel.Y+panel.PaddingY, panel.X+panel.PaddingX)

	n := panel.Height - 2 // Number of logs to display
	lenLog := len(p.State.Console.Lines)

	if lenLog == 0 {
		return
	} else if lenLog > 0 && lenLog < panel.Height-2 {
		n = lenLog
	}

	for i, log := range p.State.Console.LastN(n) {
		fmt.Printf("%s", log)
		p.State.Cursor.Move(panel.Y+panel.PaddingY+i+1, panel.X+panel.PaddingX)
	}
}

func (p *PanelView) Draw(panel *models.Panel) {
	switch panel.Format {
	case "nav":
		p.DrawNavPanel(panel)
	case "option":
		p.DrawOptionPanel(panel)
	case "description":
		p.DrawDescriptionPanel(panel)
	case "action":
		p.DrawActionPanel(panel)
	case "section":
		p.DrawSectionPanel(panel)
	case "log":
		p.DrawLogPanel(panel)
	default:
		p.DrawBox(panel, false)
	}
}

func (p *PanelView) DrawPanels(panels *[]models.Panel) {
	for _, panel := range *panels {
		p.Draw(&panel)
	}
}
