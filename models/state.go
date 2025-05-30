package models

import (
	"fmt"
	"os"

	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

type State struct {
	Console         IOConsole
	Pages           []Page
	SelectedPage    int
	SelectedSection int
	SelectedPanel   int
	SectionCursor   int
	OptionCursor    int
	ToggleOn        bool
	Cursor          Coord
	OldState        *term.State
}

func NewState() *State {
	state := &State{
		Console:         IOConsole{Lines: make([]string, 0)},
		Pages:           nil,
		SelectedPage:    0,
		SelectedSection: 0,
		SelectedPanel:   1,
		SectionCursor:   0,
		OptionCursor:    0,
		ToggleOn:        false,
		Cursor:          Coord{X: 1, Y: 1},
		OldState:        nil,
	}
	return state
}

func (s *State) Clear() {
	fmt.Print("\033[2J") // Clear screen
}

func (s *State) MoveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
	s.Cursor.X = col
	s.Cursor.Y = row
}

func (s *State) LoadSections(file string, targetPage int) {
	// Load the packages from the config file
	// This is a placeholder function, you need to implement the actual loading logic
	fmt.Println("Loading packages...")
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error opening packages file:", err)
		return
	}
	var sections []Section
	err = yaml.Unmarshal(data, &sections)
	if err != nil {
		fmt.Println("Error decoding packages file:", err)
		return
	}
	if targetPage < 0 || targetPage >= len(s.Pages) {
		fmt.Println("Error: Invalid target page index.")
		return
	}
	s.Pages[targetPage].Sections = sections
}

func (s *State) SaveSections(file string) {
	fmt.Println("Saving sections to", file)
	data, err := yaml.Marshal(s.Pages[s.SelectedPage].Sections)
	if err != nil {
		fmt.Println("Error encoding sections:", err)
		return
	}
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		fmt.Println("Error writing sections to file:", err)
		return
	}
}

func (s *State) LoadPages(file string) {
	// Load the config file and parse it into the Packages variable
	// This is a placeholder function, you need to implement the actual loading logic
	fmt.Println("Loading layout...")
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	var layout Layout
	err = yaml.Unmarshal(data, &layout)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}
	s.Pages = layout.Pages
}

func (s *State) SaveOldState() error {
	oldstate, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	s.OldState = oldstate
	return nil
}

func (s *State) RestoreOldState() error {
	if s.OldState == nil {
		return fmt.Errorf("no old state to restore")
	}
	err := term.Restore(int(os.Stdin.Fd()), s.OldState)
	if err != nil {
		return fmt.Errorf("failed to restore old state: %w", err)
	}
	s.OldState = nil // Clear the old state after RemoveSelectedOptions
	return nil
}

func (s *State) SelectSection() {
	if s.SelectedPanel == 1 {
		s.SelectedSection = s.SectionCursor
		s.OptionCursor = 0
		s.SelectedPanel = 2
	}
}

func (s *State) TogglePage() {
	if s.SelectedPage < len(s.Pages)-1 {
		s.SelectedPage++
	} else {
		s.SelectedPage = 0
	}
	s.SectionCursor = 0
	s.OptionCursor = 0
	s.SelectedPanel = 1
	s.SelectedSection = 0
}

func (s *State) MoveCursorUp() {
	if s.SelectedPanel == 1 && s.SectionCursor > 0 {
		s.SectionCursor--
	} else if s.SelectedPanel == 2 && s.OptionCursor > 0 {
		s.OptionCursor--
	}
}

func (s *State) MoveCursorDown() {
	if s.SelectedPanel == 1 && s.SectionCursor < len(s.Pages[s.SelectedPage].Sections)-1 {
		s.SectionCursor++
	} else if s.SelectedPanel == 2 && s.OptionCursor < len(s.Pages[s.SelectedPage].Sections[s.SelectedSection].Options)-1 {
		s.OptionCursor++
	}
}

func (s *State) MoveCursorLeft() {
	if s.SelectedPanel == 2 {
		s.SelectedPanel = 1
		s.OptionCursor = 0
		s.SectionCursor = 0
	}
}

func (s *State) MoveCursorRight() {
	if s.SelectedPanel == 1 {
		s.SelectedPanel = 2
		s.OptionCursor = 0
		s.SectionCursor = 0
	}
}

func (s *State) ToggleSelectOption() {
	opt := &s.Pages[s.SelectedPage].Sections[s.SelectedSection].Options[s.OptionCursor]
	if s.SelectedPanel == 2 {
		opt.Selected = !opt.Selected
	}
}

func (s *State) ToggleAllOptions() {
	s.ToggleOn = !s.ToggleOn
	for i := range s.Pages[s.SelectedPage].Sections {
		for j := range s.Pages[s.SelectedPage].Sections[i].Options {
			s.Pages[s.SelectedPage].Sections[i].Options[j].Selected = s.ToggleOn
		}
	}
}

func (s *State) RemoveSelectedOptions() {
	if s.SelectedPage != 0 {
		s.Console.Add("You can only remove packages not configuration script.")
		return
	}
	sections := s.Pages[s.SelectedPage].Sections
	if sections == nil {
		s.Console.Add("No sections available for removal.")
		return
	}
	for _, section := range sections {
		for i := len(section.Options) - 1; i >= 0; i-- {
			if section.Options[i].Selected {
				s.Console.Add(fmt.Sprintf("Removing %s...", section.Options[i].Name))
				cmd := []string{"yay", "-Rns", section.Options[i].Name, "--noconfirm", "--needed"}
				inputChan := make(chan byte, 100)
				doneChan := make(chan struct{})
				s.Console.RunCommandInPanel(cmd, inputChan, doneChan)
			}
		}
	}
}

func (s *State) InstallSelectedOptions() {
	sections := &s.Pages[s.SelectedPage].Sections
	if sections == nil {
		s.Console.Add("No sections available for installation.")
		return
	}
	for _, section := range *sections {
		for _, option := range section.Options {
			if option.Selected {
				if s.SelectedPage == 0 {
					s.Console.Add(fmt.Sprintf("Installing %s...", option.Name))
					cmd := []string{"yay", "-S", option.Name, "--noconfirm", "--needed"}
					inputChan := make(chan byte, 100)
					doneChan := make(chan struct{})
					s.Console.RunCommandInPanel(cmd, inputChan, doneChan)
				} else if s.SelectedPage == 1 && option.Scripts != "" {
					s.Console.Add(fmt.Sprintf("Executing script %s...", option.Name))
					cmd := []string{"bash", option.Scripts}
					inputChan := make(chan byte, 100)
					doneChan := make(chan struct{})
					s.Console.RunCommandInPanel(cmd, inputChan, doneChan)
				}
			}
		}
	}
}
