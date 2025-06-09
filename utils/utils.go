package utils

import (
	"strings"
)

func WrapWords(input string, maxWidth int) []string {
	var result []string
	var line string
	if len(input) <= maxWidth {
		return []string{input}
	}
	words := strings.Fields(input) // split by space, removing extra spaces

	for _, word := range words {
		if len(line)+len(word)+1 > maxWidth {
			result = append(result, strings.TrimSpace(line))
			line = word + " "
		} else {
			line += word + " "
		}
	}

	if len(line) > 0 {
		result = append(result, strings.TrimSpace(line))
	}

	return result
}

func SplitLines(input string) []string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return lines
}

func FormatLines(input string) []string {
	var result []string
	lines := SplitLines(input)
	for _, line := range lines {
		wraped := WrapWords(line, 115)
		result = append(result, wraped...)
	}
	return result
}

func TrimUntil(s string, ch rune) string {
	i := strings.IndexRune(s, ch)
	if i == -1 {
		return s
	}
	return s[i+1:]
}
