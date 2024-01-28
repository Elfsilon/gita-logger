package gita

import (
	"fmt"
	"time"
)

type Event struct {
	id         int
	message    string
	timestamp  time.Time
	filename   string
	line       int
	stackTrace string
}

func (e *Event) format(level Level) string {
	t := e.timestamp.Format("15:01:02")

	formatted := fmt.Sprintf(
		"#%d %s [%s] (%s:%d) %s ",
		e.id,
		t,
		labels[level],
		e.filename,
		e.line,
		e.message,
	)

	if e.stackTrace != "" && level >= ErrorLevel {
		formatted = fmt.Sprintf("%s\nStackTrace: %s", formatted, e.stackTrace)
	}

	return formatted
}
