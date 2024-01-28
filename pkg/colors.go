package gita

import (
	"fmt"
	"runtime"
)

type Colors struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	Gray   string
	White  string
}

func NewColors() *Colors {
	if runtime.GOOS == "windows" {
		return &Colors{
			Reset:  "",
			Red:    "",
			Green:  "",
			Yellow: "",
			Blue:   "",
			Purple: "",
			Cyan:   "",
			Gray:   "",
			White:  "",
		}
	}

	return &Colors{
		Reset:  "\033[0m",
		Red:    "\033[31m",
		Green:  "\033[32m",
		Yellow: "\033[33m",
		Blue:   "\033[34m",
		Purple: "\033[35m",
		Cyan:   "\033[36m",
		Gray:   "\033[37m",
		White:  "\033[97m",
	}
}

func (c *Colors) getColorFromLevel(level Level) string {
	var color string

	switch level {
	case InfoLevel:
		color = c.Blue
	case WarningLevel:
		color = c.Yellow
	case ErrorLevel:
		color = c.Red
	case FatalLevel:
		color = c.Red
	default:
		color = ""
	}

	return color
}

func (c *Colors) ColoredLevel(text string, level Level) string {
	color := c.getColorFromLevel(level)
	return c.Colored(text, color)
}

func (c *Colors) Colored(text string, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, c.Reset)
}
