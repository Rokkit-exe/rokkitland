package models

type Screen struct {
	InputManager InputManager
	State        *State
}

func (s *Screen) Draw() error {
	for {
		s.State.Clear()
		s.State.MoveCursor(1, 1)
		s.DrawPages()
		s.State.Pages[s.State.SelectedPage].DrawPanels(s.State)
		err := s.InputManager.RecordKeys(s.State)
		if err != nil {
			s.State.MoveCursor(30, 1)
			s.State.Log.Add("Error: " + err.Error())
			return err
		}
	}
}

func (s *Screen) DrawPages() {
	for _, page := range s.State.Pages {
		page.DrawTab(s.State, s.State.SelectedPage == page.Id)
	}
}
