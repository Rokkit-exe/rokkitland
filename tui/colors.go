package tui

type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Reset
)

func (c Color) ANSI() string {
	switch c {
	case Black:
		return "\033[30m"
	case Red:
		return "\033[31m"
	case Green:
		return "\033[32m"
	case Yellow:
		return "\033[33m"
	case Blue:
		return "\033[34m"
	case Magenta:
		return "\033[35m"
	case Cyan:
		return "\033[36m"
	case White:
		return "\033[37m"
	case Reset:
		return "\033[0m"
	default:
		return ""
	}
}
