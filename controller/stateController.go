package controller

import (
	"fmt"
	"os"

	"github.com/Rokkit-exe/rokkitland/models"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

type StateController struct {
	State             *models.State
	ConsoleController *ConsoleController
}

func NewStateController(state *models.State, consoleController *ConsoleController) *StateController {
	return &StateController{
		State:             state,
		ConsoleController: consoleController,
	}
}

func (s *StateController) LoadSections(file string, targetPage int) {
	// Load the packages from the config file
	// This is a placeholder function, you need to implement the actual loading logic
	fmt.Println("Loading packages...")
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error opening packages file:", err)
		return
	}
	var sections []models.Section
	err = yaml.Unmarshal(data, &sections)
	if err != nil {
		fmt.Println("Error decoding packages file:", err)
		return
	}
	if targetPage < 0 || targetPage >= len(s.State.Pages) {
		fmt.Println("Error: Invalid target page index.")
		return
	}
	s.State.Pages[targetPage].Sections = sections
}

func (s *StateController) SaveSections(file string) {
	fmt.Println("Saving sections to", file)
	data, err := yaml.Marshal(s.State.Pages[s.State.SelectedPage].Sections)
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

func (s *StateController) LoadPages(file string) {
	// Load the config file and parse it into the Packages variable
	// This is a placeholder function, you need to implement the actual loading logic
	fmt.Println("Loading layout...")
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	var layout models.Layout
	err = yaml.Unmarshal(data, &layout)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}
	s.State.Pages = layout.Pages
}

func (s *StateController) SaveOldState() error {
	oldstate, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	s.State.OldState = oldstate
	return nil
}

func (s *StateController) RestoreOldState() error {
	if s.State.OldState == nil {
		return fmt.Errorf("no old state to restore")
	}
	err := term.Restore(int(os.Stdin.Fd()), s.State.OldState)
	if err != nil {
		return fmt.Errorf("failed to restore old state: %w", err)
	}
	s.State.OldState = nil // Clear the old state after RemoveSelectedOptions
	return nil
}

func (s *StateController) SelectSection() {
	if s.State.SelectedPanel == 1 {
		s.State.SelectedSection = s.State.SectionCursor
		s.State.OptionCursor = 0
		s.State.SelectedPanel = 2
	}
}

func (s *StateController) TogglePage() {
	if s.State.SelectedPage < len(s.State.Pages)-1 {
		s.State.SelectedPage++
	} else {
		s.State.SelectedPage = 0
	}
	s.State.SectionCursor = 0
	s.State.OptionCursor = 0
	s.State.SelectedPanel = 1
	s.State.SelectedSection = 0
}

func (s *StateController) MoveCursorUp() {
	if s.State.SelectedPanel == 1 && s.State.SectionCursor > 0 {
		s.State.SectionCursor--
	} else if s.State.SelectedPanel == 2 && s.State.OptionCursor > 0 {
		s.State.OptionCursor--
	}
}

func (s *StateController) MoveCursorDown() {
	if s.State.SelectedPanel == 1 && s.State.SectionCursor < len(s.State.Pages[s.State.SelectedPage].Sections)-1 {
		s.State.SectionCursor++
	} else if s.State.SelectedPanel == 2 && s.State.OptionCursor < len(s.State.Pages[s.State.SelectedPage].Sections[s.State.SelectedSection].Options)-1 {
		s.State.OptionCursor++
	}
}

func (s *StateController) MoveCursorLeft() {
	if s.State.SelectedPanel == 2 {
		s.State.SelectedPanel = 1
		s.State.OptionCursor = 0
		s.State.SectionCursor = 0
	}
}

func (s *StateController) MoveCursorRight() {
	if s.State.SelectedPanel == 1 {
		s.State.SelectedPanel = 2
		s.State.OptionCursor = 0
		s.State.SectionCursor = 0
	}
}

func (s *StateController) ToggleSelectOption() {
	opt := &s.State.Pages[s.State.SelectedPage].Sections[s.State.SelectedSection].Options[s.State.OptionCursor]
	if s.State.SelectedPanel == 2 {
		opt.Selected = !opt.Selected
	}
}

func (s *StateController) ToggleAllOptions() {
	s.State.ToggleOn = !s.State.ToggleOn
	for i := range s.State.Pages[s.State.SelectedPage].Sections {
		for j := range s.State.Pages[s.State.SelectedPage].Sections[i].Options {
			s.State.Pages[s.State.SelectedPage].Sections[i].Options[j].Selected = s.State.ToggleOn
		}
	}
}

func (s *StateController) RemoveSelectedOptions() {
}

func (s *StateController) InstallSelectedOptions() {
	sections := s.State.Pages[s.State.SelectedPage].Sections

	if len(sections) == 0 {
		s.State.Console.Add("No sections available.")
		return
	}

	var cmd []string

	if s.State.SelectedPage == 0 {
		s.State.Console.Add("Installing ...")
		// cmd = []string{"yay", "-S", "--needed", "--noconfirm" option.Name}
		cmd = []string{"echo", "Hello"}
	}

	// Run command and wait for completion
	s.ConsoleController.RunCommandWithPTY(cmd)
	s.State.Console.Add("Done installing")
}
