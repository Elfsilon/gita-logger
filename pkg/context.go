package gita

import (
	"fmt"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

type Context struct {
	eventsCount        int
	attachStackTrace   bool
	col                *Colors
	displayID          bool
	displayTime        bool
	displayFileAndLine bool
}

func NewContext(
	attachStackTrace bool,
	displayID bool,
	displayTime bool,
	displayFileAndLine bool,
) *Context {
	return &Context{
		col:                NewColors(),
		eventsCount:        0,
		attachStackTrace:   attachStackTrace,
		displayID:          displayID,
		displayTime:        displayTime,
		displayFileAndLine: displayFileAndLine,
	}
}

func (c *Context) NewEventFromMessage(message string, level Level, depth int) Event {
	c.eventsCount = c.eventsCount + 1
	_, file, line, _ := runtime.Caller(depth)
	filename := path.Base(file)

	event := Event{
		ID:        c.eventsCount,
		Message:   message,
		Timestamp: time.Now(),
		Filename:  filename,
		Line:      line,
		Level:     level,
	}

	if c.attachStackTrace && level >= ErrorLevel {
		event.StackTrace = string(debug.Stack())
	}

	return event
}

func (c *Context) Format(e Event) string {
	var formatted strings.Builder

	if c.displayID {
		formatted.WriteString("#" + fmt.Sprint(e.ID) + " ")
	}

	if c.displayTime {
		t := e.Timestamp.Format("15:01:02")
		formatted.WriteString(t + " ")
	}

	coloredLevelName := c.col.ColoredLevel(labels[e.Level], e.Level)
	formatted.WriteString("[" + coloredLevelName + "] ")

	if c.displayFileAndLine {
		fileAndLine := fmt.Sprintf("%s:%d", e.Filename, e.Line)
		coloredFileAndLine := c.col.ColoredLevel(fileAndLine, e.Level)
		formatted.WriteString("(" + coloredFileAndLine + ") ")
	}

	formatted.WriteString(e.Message)

	if e.StackTrace != "" {
		coloredStack := c.col.ColoredLevel("StackTrace: "+e.StackTrace, ErrorLevel)
		formatted.WriteString("\n" + coloredStack)
	}

	return formatted.String()
}
