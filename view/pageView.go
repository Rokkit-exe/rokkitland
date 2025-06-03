package view

import (
	"fmt"
	"strings"

	"github.com/Rokkit-exe/rokkitland/models"
	"github.com/Rokkit-exe/rokkitland/tui"
)

type PageView struct {
	State *models.State
}

func NewPageView(state *models.State) *PageView {
	return &PageView{
		State: state,
	}
}

func (p *PageView) DrawTab(page *models.Page, active bool) {
	if active {
		fmt.Printf("%s", tui.Green.ANSI())
	} else {
		fmt.Printf("%s", tui.Reset.ANSI())
	}
	fmt.Printf("┌%s┐", strings.Repeat("─", len(page.Title)+2))
	p.State.Cursor.Move(p.State.Cursor.Y+1, p.State.Cursor.X)
	fmt.Printf("%s %s %s", "│", page.Title, "│")
	p.State.Cursor.Move(p.State.Cursor.Y+1, p.State.Cursor.X)
	fmt.Printf("└%s┘", strings.Repeat("─", len(page.Title)+2))
	fmt.Printf("%s", tui.Reset.ANSI())
}

func (p *PageView) DrawPages(pages *[]models.Page) {
	for i := range *pages {
		page := &(*pages)[i]
		p.DrawTab(page, p.State.SelectedPage == page.Id)
		p.State.Cursor.Move(1, len(page.Title)+5) // Move to next line
	}
}
