package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/Rokkit-exe/rokkitland/tui"
	"golang.org/x/term"
)

type Menu struct {
	Title           string
	Panels          []Panel
	Sections        []Section
	Actions         []string
	SelectedSection int
	LeftCursor      int
	RightCursor     int
	SelectedPanel   int
	OldState        *term.State
}

func (m Menu) Clear() {
	fmt.Print("\033[2J") // Clear screen
}

func (m Menu) MoveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func (m *Menu) PrintPos(index int) {
	m.MoveCursor(80, 0)
	fmt.Printf("Panel: %d, X: %d, Y: %d", m.Panels[index].Id, m.Panels[index].Pos.X, m.Panels[index].Pos.Y)
}

func (m *Menu) DrawNavPanel() {
	m.DrawPanel(0, false)
	m.PrintPos(0)
	m.MoveCursor(m.Panels[0].Pos.Y+1, m.Panels[0].Pos.X+3)
	for i, msg := range m.Panels[0].Content {
		m.MoveCursor(m.Panels[0].Pos.Y+1+i, m.Panels[0].Pos.X+3)
		fmt.Print(msg)
	}
}

func (m *Menu) DrawPanel(index int, active bool) {
	if active {
		fmt.Printf("%s", tui.Green.ANSI())
	} else {
		fmt.Printf("%s", tui.Reset.ANSI())
	}
	m.MoveCursor(m.Panels[index].Pos.Y, m.Panels[index].Pos.X)
	fmt.Printf("┌─%s%s┐", m.Panels[index].Title, strings.Repeat("─", m.Panels[index].Dimensions.Width-len(m.Panels[index].Title)-2))
	for i := 1; i < m.Panels[index].Dimensions.Height-1; i++ {
		m.MoveCursor(m.Panels[index].Pos.Y+i, m.Panels[index].Pos.X)
		fmt.Print("│" + strings.Repeat(" ", m.Panels[index].Dimensions.Width-1) + "│")
	}
	m.MoveCursor(m.Panels[index].Pos.Y+m.Panels[index].Dimensions.Height-1, m.Panels[index].Pos.X)
	fmt.Printf("└%s┘", strings.Repeat("─", m.Panels[index].Dimensions.Width-1))
	fmt.Printf("%s", tui.Reset.ANSI())
}

func (m *Menu) DrawSectionPanel() {
	if m.SelectedPanel == 1 {
		m.DrawPanel(1, true)
	} else {
		m.DrawPanel(1, false)
	}
	m.MoveCursor(m.Panels[1].Pos.Y+1, m.Panels[1].Pos.X+3)
	// left Panels[1]
	for i, opt := range m.Sections {
		cursorPrefix := "  "
		if i == m.LeftCursor && m.SelectedPanel == 1 {
			m.SelectedSection = i
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		m.MoveCursor(m.Panels[1].Pos.Y+1+i, m.Panels[1].Pos.X+3)
		fmt.Printf("%s%s", cursorPrefix, opt.Title)
	}
}

func (m *Menu) DrawActionsPanel() {
	if m.SelectedPanel == 3 {
		m.DrawPanel(3, true)
	} else {
		m.DrawPanel(3, false)
	}
	m.MoveCursor(m.Panels[3].Pos.Y+1, m.Panels[3].Pos.X+3)
	for i, opt := range m.Panels[3].Content {
		cursorPrefix := "  "
		if i == m.LeftCursor && m.SelectedPanel == 3 {
			m.SelectedSection = i
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		m.MoveCursor(m.Panels[3].Pos.Y+1+i, m.Panels[3].Pos.X+3)
		fmt.Printf("%s%s", cursorPrefix, opt)
	}
}

func (m *Menu) DrawOptionsPanel() {
	if m.SelectedPanel == 2 {
		m.DrawPanel(2, true)
	} else {
		m.DrawPanel(2, false)
	}
	m.MoveCursor(m.Panels[2].Pos.Y+1, m.Panels[2].Pos.X+3)
	// right Panel
	for i, opt := range m.Sections[m.SelectedSection].Packages {
		prefix := "[ ]"
		if opt.Selected {
			prefix = fmt.Sprintf("[%sx%s]", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		cursorPrefix := "  "
		if i == m.RightCursor && m.SelectedPanel == 2 {
			cursorPrefix = fmt.Sprintf("%s>%s ", tui.Green.ANSI(), tui.Reset.ANSI())
		}
		m.MoveCursor(m.Panels[2].Pos.Y+1+i, m.Panels[2].Pos.X+3)
		fmt.Printf("%s%s %s", cursorPrefix, prefix, opt.Name)
	}
}

func (m *Menu) RecordKeys() error {
	buf := make([]byte, 3)
	os.Stdin.Read(buf)

	if buf[0] == tui.Enter { // Enter
		if m.SelectedPanel == 1 {
			m.SelectedSection = m.LeftCursor
			m.RightCursor = 0
			m.SelectedPanel = 2
		}
	}
	if buf[0] == tui.Escape1 && buf[1] == tui.Escape2 { // Arrow keys
		switch buf[2] {
		case tui.Up: // Up
			if m.SelectedPanel == 1 && m.LeftCursor > 0 {
				m.LeftCursor--
			} else if m.SelectedPanel == 2 && m.RightCursor > 0 {
				m.RightCursor--
			}
		case tui.Down: // Down
			if m.SelectedPanel == 1 && m.LeftCursor < len(m.Sections)-1 {
				m.LeftCursor++
			} else if m.SelectedPanel == 2 && m.RightCursor < len(m.Sections[m.SelectedSection].Packages)-1 {
				m.RightCursor++
			}
		case tui.Left: // Left
			if m.SelectedPanel == 2 {
				m.SelectedPanel = 1
			}
		case tui.Right: // Right
			if m.SelectedPanel == 1 {
				m.SelectedPanel = 2
			}
		}
	}
	if buf[0] == tui.Space && m.SelectedPanel == 2 {
		if m.Sections[m.SelectedSection].Packages[m.RightCursor].Selected {
			m.Sections[m.SelectedSection].Packages[m.RightCursor].Selected = false
		} else {
			m.Sections[m.SelectedSection].Packages[m.RightCursor].Selected = true
		}
	}

	if buf[0] == tui.Quit { // Escape1
		m.MoveCursor(90, 0)
		fmt.Println("Exiting...")
		time.Sleep(1 * time.Second)
		m.Clear()
		term.Restore(int(syscall.Stdin), m.OldState)
		os.Exit(0)
	}
	if buf[0] == tui.Tab { // Tab
		if m.SelectedPanel != 2 {
			m.SelectedPanel = 2
		} else {
			m.SelectedPanel = 0
		}
	}
	return nil
}

func (m *Menu) DrawMenu() error {
	oldstate, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		return err
	}
	m.OldState = oldstate
	defer term.Restore(int(syscall.Stdin), m.OldState)

	for {
		m.Clear()
		m.MoveCursor(1, 1)
		m.DrawNavPanel()
		m.DrawSectionPanel()
		m.DrawOptionsPanel()
		m.DrawActionsPanel()

		err := m.RecordKeys()
		if err != nil {
			m.MoveCursor(90, 0)
			fmt.Println("Error:", err)
			return err
		}
	}
}

func (m *Menu) LoadConfig() {
	// Load the config file and parse it into the Packages variable
	// This is a placeholder function, you need to implement the actual loading logic
	fmt.Println("Loading config...")
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer configFile.Close()
	var Panels []Panel
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&Panels)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}
	m.Panels = Panels
}

func (m *Menu) LoadSections() {
	// Load the packages from the config file
	// This is a placeholder function, you need to implement the actual loading logic
	fmt.Println("Loading packages...")
	configFile, err := os.Open("packages.json")
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
	m.Sections = sections
}
