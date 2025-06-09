package models

import (
	"slices"
	"sync"

	"github.com/Rokkit-exe/rokkitland/tui"
	"github.com/Rokkit-exe/rokkitland/utils"
)

type Console struct {
	Lines []string
	Mu    sync.Mutex
}

func NewConsole() *Console {
	return &Console{
		Lines: make([]string, 0),
		Mu:    sync.Mutex{},
	}
}

func (c *Console) Add(line string, color tui.Color) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if len(c.Lines) > 200 {
		c.Lines = c.Lines[len(c.Lines)-200:]
	}
	lines := utils.WrapWords(line, 115)
	for _, l := range lines {
		c.Lines = append(c.Lines, tui.Colorize(l, color))
	}
}

func (c *Console) GetLines() []string {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return slices.Clone(c.Lines)
}

func (c *Console) LastN(n int) []string {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if len(c.Lines) < n {
		return c.Lines
	}
	return c.Lines[len(c.Lines)-n:]
}
