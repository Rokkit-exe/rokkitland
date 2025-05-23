package models

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

type Menu struct {
	Title        string
	Pages        []Page
	InputManager InputManager
	Log          Log
	State        State
}

func (m *Menu) DrawMenu() error {
	err := m.State.SaveOldState()
	if err != nil {
		m.Log.Add("Error saving old state: " + err.Error())
	}
	defer term.Restore(int(syscall.Stdin), m.State.OldState)

	for {
		m.State.Clear()
		m.State.MoveCursor(0, 0)
		for _, page := range m.Pages {
			page.DrawTab(&m.State, true)
		}
		// m.Pages[m.State.SelectedPage].DrawPanels(&m.State, m.InputManager)
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
