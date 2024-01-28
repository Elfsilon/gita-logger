package gita

import (
	"path"
	"runtime"
	"runtime/debug"
	"time"
)

type Context struct {
	eventsCount      int
	attachStackTrace bool
}

func NewContext(attachStackTrace bool) *Context {
	return &Context{
		eventsCount:      0,
		attachStackTrace: attachStackTrace,
	}
}

func (c *Context) NewEventFromMessage(message string) *Event {
	c.eventsCount = c.eventsCount + 1
	_, file, line, _ := runtime.Caller(3)
	filename := path.Base(file)

	event := &Event{
		id:        c.eventsCount,
		message:   message,
		timestamp: time.Now(),
		filename:  filename,
		line:      line,
	}

	if c.attachStackTrace {
		event.stackTrace = string(debug.Stack())
	}

	return event
}
