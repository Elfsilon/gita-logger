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
		id:        c.eventsCount,
		message:   message,
		timestamp: time.Now(),
		filename:  filename,
		line:      line,
		level:     level,
	}

	if c.attachStackTrace {
		event.stackTrace = string(debug.Stack())
	}

	return event
}

func (c *Context) Format(e Event) string {
	var formatted strings.Builder

	if c.displayID {
		formatted.WriteString("#" + fmt.Sprint(e.id) + " ")
	}

	if c.displayTime {
		t := e.timestamp.Format("15:01:02")
		formatted.WriteString(t + " ")
	}

	coloredLevelName := c.col.ColoredLevel(labels[e.level], e.level)
	formatted.WriteString("[" + coloredLevelName + "] ")

	if c.displayFileAndLine {
		fileAndLine := fmt.Sprintf("%s:%d", e.filename, e.line)
		coloredFileAndLine := c.col.ColoredLevel(fileAndLine, e.level)
		formatted.WriteString("(" + coloredFileAndLine + ") ")
	}

	formatted.WriteString(e.message)

	if e.stackTrace != "" && e.level >= ErrorLevel {
		coloredStack := c.col.ColoredLevel("StackTrace: "+e.stackTrace, ErrorLevel)
		formatted.WriteString("\n" + coloredStack)
	}

	return formatted.String()
}
