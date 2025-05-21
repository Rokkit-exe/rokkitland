package main

import (
	"fmt"
	"time"

	"github.com/Rokkit-exe/rokkitland/art"
	"github.com/Rokkit-exe/rokkitland/models"
)

func main() {
	fmt.Printf(art.LOGO)
	fmt.Println("Welcome to Rokkitland! Time to create the best Arch Linux experience.")
	time.Sleep(2 * time.Second)
	mainMenu := models.Menu{
		Title: "Rokkitland Installer",
		State: models.State{
			SelectedPage:    0,
			SelectedSection: 0,
			SelectedPanel:   1,
			SectionCursor:   0,
			OptionCursor:    0,
			ActionCursor:    0,
			Cursor:          models.Coord{0, 0},
		},
	}

	mainMenu.State.LoadPages()
	mainMenu.State.LoadSections()

	if mainMenu.State.Pages == nil {
		fmt.Println("Error: No pages found.")
		return
	}

	if mainMenu.State.Sections == nil {
		fmt.Println("Error: No sections found.")
		return
	}

	time.Sleep(5 * time.Second)
	mainMenu.State.Clear()

	err := mainMenu.DrawMenu()
	if err != nil {
		mainMenu.State.Clear()
		fmt.Println("Error:", err)
		return
	}
}
