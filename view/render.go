package view

import (
	"github.com/Rokkit-exe/rokkitland/models"
)

type Renderer struct {
	State     *models.State
	PageView  *PageView
	PanelView *PanelView
}

func NewRenderer(state *models.State) *Renderer {
	return &Renderer{
		State:     state,
		PageView:  NewPageView(state),
		PanelView: NewPanelView(state),
	}
}

func (r *Renderer) Render() {
	r.State.ClearScreen()
	r.State.Cursor.Move(1, 1)
	r.PageView.DrawPages(&r.State.Pages)
	r.PanelView.DrawPanels(&r.State.Pages[r.State.SelectedPage].Panels)
}
