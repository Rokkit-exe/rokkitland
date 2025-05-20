package packages

import (
	"fmt"
	"os/exec"
	"strings"
)

type Package struct {
	Name        string
	Selected    bool
	Description string
}

func (p Package) String() string {
	return fmt.Sprintf("%s | %s", p.Name, p.Description)
}

func (p *Package) GetDescription() {
	cmd := exec.Command("pacman", "-Qi", p.Name, "|", "grep", "Description")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	description := strings.TrimPrefix(string(output), "Description     : ")
	p.Description = description
}

var Sections = []string{
	"Desktop Environment",
	"Styling",
	"System Utilities",
	"Development",
	"Media",
}

var Packages = [][][]string{
	{
		{"hyprland", "true"},
		{"hyprlock", "true"},
		{"hyprpicker", "true"},
		{"hyprpaper", "true"},
		{"hypridle", "true"},
		{"hyprshot", "true"},
		{"swaync", "true"},
		{"waybar", "true"},
		{"wlogout", "true"},
		{"wofi", "true"},
		{"firefox", "true"},
		{"thunar", "true"},
		{"kitty", "true"},
		{"zsh", "true"},
		{"gwenview", "true"},
	},
	{
		{"adw-gtk-theme", "true"},
		{"papirus-icon-theme", "true"},
		{"catppuccin-gtk-theme-mocha", "true"},
		{"catppuccin-sddm-theme-mocha", "true"},
		{"ttf-font-awesome", "true"},
		{"ttf-ubuntu-mono-nerd", "true"},
		{"ttf-ubuntu-nerd", "true"},
		{"simple-sddm-theme-2-git", "true"},
		{"simple-sddm-theme-git", "true"},
	},
	{
		{"wget", "true"},
		{"curl", "true"},
		{"neofetch", "true"},
		{"nvtop", "true"},
		{"htop", "true"},
		{"radeontop", "true"},
		{"nwg-look", "true"},
		{"man-db", "true"},
		{"man-pages", "true"},
		{"pamixer", "true"},
		{"pavucontrol", "true"},
		{"blueman", "true"},
		{"bluez", "true"},
		{"bluez-utils", "true"},
	},
	{
		{"git", "true"},
		{"neovim", "true"},
		{"docker", "true"},
		{"mariadb", "true"},
		{"go", "true"},
		{"rust", "true"},
		{"nodejs", "true"},
		{"postman-git", "true"},
		{"beekeeper-studio-bin", "true"},
		{"bun-bin", "true"},
	},
	{
		{"discord", "true"},
		{"qbittorrent", "true"},
		{"thunderbird", "true"},
		{"steam", "true"},
		{"obs-studio", "true"},
		{"obsidian", "true"},
		{"vlc", "true"},
		{"etcher-bin", "true"},
	},
}
