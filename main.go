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
	fmt.Printf(art.LOGO)
	fmt.Println("Welcome to Rokkitland! Time to create the best Arch Linux experience.")

	time.Sleep(2 * time.Second)

	state := models.State{
		SelectedPage:    0,
		SelectedSection: 0,
		SelectedPanel:   1,
		SectionCursor:   0,
		OptionCursor:    0,
		ActionCursor:    0,
		Cursor:          models.Coord{X: 1, Y: 1},
	}

	state.LoadPages()
	state.LoadSections()

	if state.Pages == nil {
		fmt.Println("Error: No pages found.")
		return
	}

	if state.Sections == nil {
		fmt.Println("Error: No sections found.")
		return
	}
	fmt.Println("Pages loaded successfully.")

	time.Sleep(5 * time.Second)

	err := state.SaveOldState()
	if err != nil {
		state.Log.Add("Error saving old state: " + err.Error())
	}
	defer term.Restore(int(syscall.Stdin), state.OldState)

	screen := models.Screen{
		InputManager: models.InputManager{},
		State:        &state,
	}

	err = screen.Draw()
	if err != nil {
		state.Log.Add("Error drawing menu: " + err.Error())
		fmt.Println("Error:", err)
		return
	}
}
