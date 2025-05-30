package models

type Section struct {
	Title   string   `json:"title"`
	Options []Option `json:"packages"`
}

func (s *Section) UpdateDescription() {
	for i := range s.Options {
		s.Options[i].Description = s.Options[i].GetDescription()
	}
}
