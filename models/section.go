package models

type Section struct {
	Title    string    `json:"title"`
	Packages []Package `json:"packages"`
}
