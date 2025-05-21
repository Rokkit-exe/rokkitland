package models

type Section struct {
	Title   string   `json:"title"`
	Options []Option `json:"packages"`
}
