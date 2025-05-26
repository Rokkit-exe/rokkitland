package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Menu struct {
	Title        string
	Pages        []Page
	InputManager InputManager
	Log          Log
	State        State
}

func (m *Menu) DrawMenu() error {
	for {
		m.State.Clear()
		m.State.MoveCursor(1, 1)
		m.DrawPages()
		m.Pages[m.State.SelectedPage].DrawPanels(&m.State)
		err := m.InputManager.RecordKeys()
		if err != nil {
			m.State.MoveCursor(60, 1)
			fmt.Println("Error:", err)
			return err
		}
	}
}

func (m *Menu) DrawPages() {
	for _, page := range m.Pages {
		page.DrawTab(&m.State, m.State.SelectedPage == page.Id)
	}
}

func (m *Menu) LoadPages() {
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
	m.Pages = pages
}
