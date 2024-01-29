package gita

import (
	"path"
	"runtime"
	"runtime/debug"
	"time"
)

type Context struct {
	eventsCount int
}

func NewContext() *Context {
	return &Context{
		eventsCount: 0,
	}
}

func (c *Context) NewEventFromMessage(
	message string,
	level Level,
	attachStackTrace bool,
	depth int,
) Event {
	c.eventsCount = c.eventsCount + 1
	filename, line := c.getFilenameAndLine(depth + 1)

	event := Event{
		ID:        c.eventsCount,
		Message:   message,
		Timestamp: time.Now(),
		Filename:  filename,
		Line:      line,
		Level:     level,
	}

	if attachStackTrace && level >= ErrorLevel {
		event.StackTrace = string(debug.Stack())
	}

	return event
}

func (c *Context) getFilenameAndLine(depth int) (string, int) {
	_, file, line, _ := runtime.Caller(depth)
	filename := path.Base(file)

	return filename, line
}
