package models

type Section struct {
	Title   string   `yaml:"title"`
	Options []Option `yaml:"options"`
}

func (s *Section) UpdateDescription() {
	for i := range s.Options {
		s.Options[i].Description = s.Options[i].GetDescription()
	}
}
