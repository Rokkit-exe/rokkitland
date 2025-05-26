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
	mainMenu := models.Menu{
		Title: "Rokkitland Installer",
		State: models.State{
			SelectedPage:    0,
			SelectedSection: 0,
			SelectedPanel:   1,
			SectionCursor:   0,
			OptionCursor:    0,
			ActionCursor:    0,
			Cursor:          models.Coord{X: 0, Y: 0},
		},
	}

	mainMenu.LoadPages()
	mainMenu.State.LoadSections()

	if mainMenu.Pages == nil {
		fmt.Println("Error: No pages found.")
		return
	}

	if mainMenu.State.Sections == nil {
		fmt.Println("Error: No sections found.")
		return
	}
	fmt.Println("Pages loaded successfully.")
	pages := fmt.Sprintf("pages found: %d", len(mainMenu.Pages))
	pageTitle := fmt.Sprintf("page title: %s", mainMenu.Pages[0].Title)
	panels := fmt.Sprintf("panels found: %d", len(mainMenu.Pages[0].Panels))
	panelWidth := fmt.Sprintf("panel width: %d", mainMenu.Pages[0].Panels[0].Width)
	panelHeight := fmt.Sprintf("panel height: %d", mainMenu.Pages[0].Panels[0].Height)
	panelTitle := fmt.Sprintf("panel title: %s", mainMenu.Pages[0].Panels[0].Title)
	sections := fmt.Sprintf("sections found: %d", len(mainMenu.State.Sections))

	fmt.Println(pages)
	fmt.Println(pageTitle)
	fmt.Println(panels)
	fmt.Println(panelWidth)
	fmt.Println(panelHeight)
	fmt.Println(panelTitle)
	fmt.Println(sections)

	time.Sleep(5 * time.Second)

	err := mainMenu.State.SaveOldState()
	if err != nil {
		mainMenu.Log.Add("Error saving old state: " + err.Error())
	}
	defer term.Restore(int(syscall.Stdin), mainMenu.State.OldState)
	err = mainMenu.DrawMenu()
	if err != nil {
		mainMenu.Log.Add("Error drawing menu: " + err.Error())
		fmt.Println("Error:", err)
		return
	}
}
