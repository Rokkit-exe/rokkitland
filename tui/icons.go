package tui

type Icons int

const (
	Check Icons = iota
	Cross
	Info
	Warning
	Question
	Success
	Failure
	Loading
	ArrowUp
	ArrowDown
	ArrowLeft
	ArrowRight
	SpaceBar
	Return
	Tabulation
)

func (i Icons) ANSI() string {
	switch i {
	case Check:
		return "\u2713" // Check mark
	case Cross:
		return "\u2718" // Cross mark
	case Info:
		return "\u2139" // Information symbol
	case Warning:
		return "\u26A0" // Warning sign
	case Question:
		return "\u2753" // Question mark
	case Success:
		return "\u2705" // Green check mark
	case Failure:
		return "\u274C" // Cross mark
	case Loading:
		return "\u231B" // Hourglass symbol
	case ArrowUp:
		return "↑"
	case ArrowDown:
		return "↓"
	case ArrowLeft:
		return "←"
	case ArrowRight:
		return "→"
	case SpaceBar:
		return "␣"
	case Return:
		return "↵"
	case Tabulation:
		return "⇥"
	default:
		return ""
	}
}
