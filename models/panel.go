package models

type position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Panel struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	Format     string     `json:"format"`
	Content    []string   `json:"content"`
	Pos        position   `json:"pos"`
	Dimensions dimensions `json:"dimensions"`
}
