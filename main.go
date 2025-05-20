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
		Title:           "Rokkitland Installer",
		SelectedSection: 0,
		LeftCursor:      0,
		RightCursor:     0,
		SelectedPanel:   0,
	}

	mainMenu.LoadConfig()
	mainMenu.LoadSections()

	if mainMenu.Panels == nil {
		fmt.Println("Error: No panels found.")
		return
	}

	if mainMenu.Sections == nil {
		fmt.Println("Error: No sections found.")
		return
	}

	for _, panel := range mainMenu.Panels {
		fmt.Println(panel.Id)
		fmt.Println(panel.Title)
		fmt.Println(panel.Content)
		fmt.Println(panel.Pos)
		fmt.Println(panel.Dimensions)
	}

	time.Sleep(5 * time.Second)
	mainMenu.Clear()

	err := mainMenu.DrawMenu()
	if err != nil {
		mainMenu.Clear()
		fmt.Println("Error:", err)
		return
	}
}
