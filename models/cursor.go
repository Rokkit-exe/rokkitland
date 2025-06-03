package models

import "fmt"

type Cursor struct {
	X int
	Y int
}

func (c *Cursor) Move(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
	c.X = col
	c.Y = row
}

func (c *Cursor) Get() (int, int) {
	return c.X, c.Y
}
