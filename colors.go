package gita

import (
	"strings"
)

type TextColor string
type TextBgColor string
type TextEmphasis string

const (
	Reset string = "\x1b[0m"

	Black        TextColor = "\x1b[30m"
	DarkGray     TextColor = "\x1b[30;1m"
	Red          TextColor = "\x1b[31m"
	VividRed     TextColor = "\x1b[31;1m"
	Green        TextColor = "\x1b[32m"
	VividGreen   TextColor = "\x1b[32;1m"
	Yellow       TextColor = "\x1b[33m"
	VividYellow  TextColor = "\x1b[33;1m"
	Blue         TextColor = "\x1b[34m"
	VividBlue    TextColor = "\x1b[34;1m"
	Magenta      TextColor = "\x1b[35m"
	VividMagenta TextColor = "\x1b[35;1m"
	Cyan         TextColor = "\x1b[36m"
	VividCyan    TextColor = "\x1b[36;1m"
	Gray         TextColor = "\x1b[37m"
	White        TextColor = "\x1b[37;1m"

	BlackBG   TextBgColor = "\x1b[40m"
	RedBG     TextBgColor = "\x1b[41m"
	GreenBG   TextBgColor = "\x1b[42m"
	YellowBG  TextBgColor = "\x1b[43m"
	BlueBG    TextBgColor = "\x1b[44m"
	MagentaBG TextBgColor = "\x1b[45m"
	CyanBG    TextBgColor = "\x1b[46m"
	GrayBG    TextBgColor = "\x1b[47m"

	Bold      TextEmphasis = "\x1b[1m"
	Italic    TextEmphasis = "\x1b[3m"
	Underline TextEmphasis = "\x1b[4m"
	Blink     TextEmphasis = "\x1b[5m"
	Inverse   TextEmphasis = "\x1b[7m"
)

type Text string

func (t Text) Build(style *TextStyle) string {
	var b strings.Builder

	if style.Color != "" {
		b.WriteString(string(style.Color))
	}

	if style.BGColor != "" {
		b.WriteString(string(style.BGColor))
	}

	if style.Style != "" {
		b.WriteString(string(style.Style))
	}

	b.WriteString(string(t))
	b.WriteString(Reset)

	return b.String()
}

type TextStyle struct {
	Style   TextEmphasis
	Color   TextColor
	BGColor TextBgColor
}

func (t *TextStyle) apply(src *string) {
	if t != nil {
		*src = Text(*src).Build(t)
	}
}
