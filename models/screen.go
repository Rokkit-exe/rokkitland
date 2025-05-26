package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Screen struct {
	Pages        []Page
	InputManager InputManager
	State        State
}

func (s *Screen) Draw() error {
	for {
		s.State.Clear()
		s.State.MoveCursor(1, 1)
		s.DrawPages()
		s.Pages[s.State.SelectedPage].DrawPanels(&s.State)
		err := s.InputManager.RecordKeys(&s.State)
		if err != nil {
			s.State.MoveCursor(30, 1)
			s.State.Log.Add("Error: " + err.Error())
			return err
		}
	}
}

func (s *Screen) DrawPages() {
	for _, page := range s.Pages {
		page.DrawTab(&s.State, s.State.SelectedPage == page.Id)
	}
}

func (s *Screen) LoadPages() {
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
