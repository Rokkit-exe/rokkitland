package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/Rokkit-exe/rokkitland/art"
	"github.com/Rokkit-exe/rokkitland/models"
	"golang.org/x/term"
)

func main() {
	layoutFile := "config/layout.yml"
	packagesFile := "config/packages.yml"
	configFile := "config/config.yml"
	fmt.Printf(art.LOGO)
	fmt.Println("Welcome to Rokkitland! Time to create the best Arch Linux experience.")

	time.Sleep(2 * time.Second)

	state := models.NewState()
	state.LoadPages(layoutFile)
	state.LoadSections(packagesFile, 0)
	state.LoadSections(configFile, 1)

	if state.Pages == nil {
		fmt.Println("Error: No pages found.")
		return
	} else {
		fmt.Println("Pages loaded successfully.")
	}

	if state.Pages[state.SelectedPage].Sections == nil {
		fmt.Println("Error: No sections found.")
		return
	} else {
		fmt.Println("Sections loaded successfully.")
	}

	time.Sleep(5 * time.Second)

	err := state.SaveOldState()
	if err != nil {
		state.Console.Add("Error saving old state: " + err.Error())
	}
	defer term.Restore(int(syscall.Stdin), state.OldState)

	screen := models.Screen{
		InputManager: models.InputManager{},
		State:        state,
	}

	err = screen.Draw()
	if err != nil {
		state.Console.Add("Error drawing menu: " + err.Error())
		fmt.Println("Error:", err)
		return
	}
}
