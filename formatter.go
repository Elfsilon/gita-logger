package gita

import (
	"fmt"
	"maps"
	"strings"
)

type LevelStyle map[Level]TextStyle

type Formatter struct {
	DisplayID         bool
	DisplayTimestamp  bool
	DisplayLocation   bool
	DisplayStackTrace bool
	TimestampFormat   string
	LevelStyle        LevelStyle
	IDStyle           *TextStyle
	TimestampStyle    *TextStyle
	MessageStyle      *TextStyle
	StackTraceStyle   *TextStyle
}

func NewDefaultFormatter() *Formatter {
	return &Formatter{
		DisplayStackTrace: true,
		DisplayLocation:   true,
		DisplayID:         false,
		DisplayTimestamp:  false,
		TimestampFormat:   "15:01:02",
		IDStyle:           nil,
		TimestampStyle:    nil,
		MessageStyle:      nil,
		LevelStyle: LevelStyle{
			InfoLevel: {
				Color: Blue,
			},
			WarningLevel: {
				Color: Yellow,
			},
			ErrorLevel: {
				Style: Bold,
				Color: VividRed,
			},
			FatalLevel: {
				Style: Bold,
				Color: VividRed,
			},
		},
		StackTraceStyle: &TextStyle{
			Color: Red,
		},
	}
}

func (f *Formatter) merge(other *Formatter) {
	f.DisplayID = other.DisplayID
	f.DisplayLocation = other.DisplayLocation
	f.DisplayStackTrace = other.DisplayStackTrace
	f.DisplayTimestamp = other.DisplayTimestamp

	if other.TimestampFormat != "" {
		f.TimestampFormat = other.TimestampFormat
	}

	if other.IDStyle != nil {
		f.IDStyle = other.IDStyle
	}

	if other.TimestampStyle != nil {
		f.TimestampStyle = other.TimestampStyle
	}

	if other.LevelStyle != nil {
		maps.Copy(f.LevelStyle, other.LevelStyle)
	}

	if other.MessageStyle != nil {
		f.MessageStyle = other.MessageStyle
	}

	if other.StackTraceStyle != nil {
		f.StackTraceStyle = other.StackTraceStyle
	}
}

func (f *Formatter) writeStyled(b *strings.Builder, value string, style *TextStyle) {
	style.apply(&value)
	b.WriteString(value)
}

func (f *Formatter) Format(e Event) string {
	var b strings.Builder

	if f.DisplayID {
		f.writeStyled(&b, fmt.Sprintf("#%d ", e.ID), f.IDStyle)
	}

	if f.DisplayTimestamp {
		f.writeStyled(&b, e.Timestamp.Format(f.TimestampFormat)+" ", f.TimestampStyle)
	}

	levelName := levelNames[e.Level]
	var levelStyle TextStyle
	if f.LevelStyle != nil {
		levelStyle = f.LevelStyle[e.Level]
	}
	levelStyle.apply(&levelName)
	b.WriteString("[" + levelName + "] ")

	if f.DisplayLocation {
		f.writeStyled(&b, fmt.Sprintf("(%s:%d) ", e.Filename, e.Line), &levelStyle)
	}

	f.writeStyled(&b, e.Message, f.MessageStyle)

	if e.StackTrace != "" {
		f.writeStyled(&b, "\nStackTrace: "+e.StackTrace, f.StackTraceStyle)
	}

	return b.String()
}
