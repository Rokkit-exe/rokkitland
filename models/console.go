package models

import (
	"sync"
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

func (c *Console) Add(line string) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Lines = append(c.Lines, line)
	if len(c.Lines) > 200 {
		c.Lines = c.Lines[len(c.Lines)-200:]
	}
}

func (c *Console) GetLines() []string {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return append([]string(nil), c.Lines...)
}

func (c *Console) LastN(n int) []string {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if len(c.Lines) < n {
		return c.Lines
	}
	return c.Lines[len(c.Lines)-n:]
}
