package tui

import (
	"fmt"
)

var NavMessages = []string{}

var ActionsMessages = []string{
	fmt.Sprintf("%s %s %s", "Press", Return.ANSI(), "to select an option"),
}

func Message(s string, icons string, color string) {
	fmt.Printf("%s%s %s\033[0m\n", color, icons, Reset.ANSI())
}
