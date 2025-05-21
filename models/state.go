package models

import (
	"encoding/json"
	"fmt"
	"golang.org/x/term"
	"os"
)

type State struct {
	Log             Log
	Pages           []Page
	Sections        []Section
	SelectedPage    int
	SelectedSection int
	SelectedPanel   int
	SectionCursor   int
	OptionCursor    int
	ActionCursor    int
	Cursor          Coord
	OldState        *term.State
}

func (s State) Clear() {
	fmt.Print("\033[2J") // Clear screen
}

func (s State) MoveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
	s.Cursor.X = col
	s.Cursor.Y = row
}

func (s *State) ToggleSelectedPanel() {
	if s.SelectedPanel < 3 {
		s.SelectedPanel++
	} else {
		s.SelectedPanel = 1
	}
}

func (s *State) SelectSection() {
	if s.SelectedPanel == 1 {
		s.SelectedSection = s.SectionCursor
		s.OptionCursor = 0
		s.SelectedPanel = 2
	}
}

func (s *State) SelectAction() {
	//if s.SelectedPanel == 3 {
	//	err := s.Actions[s.ActionCursor].Exec()
	//	if err != nil {
	//		s.Log.Add(err.Error())
	//	}
	//}
}

func (s *State) MoveCursorUp() {
	if s.SelectedPanel == 1 && s.SectionCursor > 0 {
		s.SectionCursor--
	} else if s.SelectedPanel == 2 && s.OptionCursor > 0 {
		s.OptionCursor--
	}
}

func (s *State) MoveCursorDown() {
	if s.SelectedPanel == 1 && s.SectionCursor < len(s.Sections)-1 {
		s.SectionCursor++
	} else if s.SelectedPanel == 2 && s.OptionCursor < len(s.Sections[s.SelectedSection].Options)-1 {
		s.OptionCursor++
	}
}

func (s *State) MoveCursorLeft() {
	if s.SelectedPanel == 2 {
		s.SelectedPanel = 1
	}
}

func (s *State) MoveCursorRight() {
	if s.SelectedPanel == 1 {
		s.SelectedPanel = 2
	}
}

func (s *State) ToggleSelectOption() {
	if s.SelectedPanel == 2 {
		if s.Sections[s.SelectedSection].Options[s.OptionCursor].Selected {
			s.Sections[s.SelectedSection].Options[s.OptionCursor].Selected = false
		} else {
			s.Sections[s.SelectedSection].Options[s.OptionCursor].Selected = true
		}
	}
}

func (s *State) LoadSections() {
	// Load the packages from the config file
	// This is a placeholder function, you need to implement the actual loading logic
	fmt.Println("Loading packages...")
	configFile, err := os.Open("config/options.json")
	if err != nil {
		fmt.Println("Error opening packages file:", err)
		return
	}
	defer configFile.Close()
	var sections []Section
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&sections)
	if err != nil {
		fmt.Println("Error decoding packages file:", err)
		return
	}
	s.Sections = sections
}

func (s *State) LoadPages() {
	// Load the config file and parse it into the Packages variable
	// This is a placeholder function, you need to implement the actual loading logic
	fmt.Println("Loading config...")
	configFile, err := os.Open("config/pages.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer configFile.Close()
	var pages []Page
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&pages)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}
	s.Pages = pages
}
