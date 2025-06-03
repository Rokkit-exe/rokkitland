package models

type Page struct {
	Id       int       `yaml:"id"`
	Title    string    `yaml:"title"`
	Panels   []Panel   `yaml:"panels"`
	Sections []Section `yaml:"sections,omitempty"`
}
