package gita

import (
	"path"
	"runtime"
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

func (c *Context) NewEventFromMessage(message string) *Event {
	c.eventsCount = c.eventsCount + 1
	_, file, line, _ := runtime.Caller(3)
	filename := path.Base(file)

	return &Event{
		id:        c.eventsCount,
		message:   message,
		timestamp: time.Now(),
		filename:  filename,
		line:      line,
	}
}
