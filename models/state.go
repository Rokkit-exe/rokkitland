package models

import (
	"fmt"
	"sync"

	"golang.org/x/term"
)

type State struct {
	Console          Console
	Pages            []Page
	SelectedPage     int
	SelectedSection  int
	SelectedPanel    int
	SectionCursor    int
	OptionCursor     int
	ToggleOn         bool
	Cursor           Cursor
	IsCommandRunning bool
	CommandInputChan chan []byte
	CommandQueue     chan QueuedCommand
	OldState         *term.State
	Mu               sync.Mutex
}

func NewState() *State {
	state := &State{
		Console:          Console{Lines: make([]string, 0)},
		Pages:            nil,
		SelectedPage:     0,
		SelectedSection:  0,
		SelectedPanel:    1,
		SectionCursor:    0,
		OptionCursor:     0,
		ToggleOn:         false,
		Cursor:           Cursor{X: 1, Y: 1},
		OldState:         nil,
		IsCommandRunning: false,
	}
	return state
}

func (s *State) ClearScreen() {
	fmt.Print("\033[2J") // Clear screen
}

func (s *State) SetIsCommandRunning(running bool) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.IsCommandRunning = running
}

func (s *State) GetIsCommandRunning() bool {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	return s.IsCommandRunning
}

func (s *State) CreateCommandInputChan() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	if s.CommandInputChan == nil {
		s.CommandInputChan = make(chan []byte, 100) // Initialize the command input channel
	}
}
