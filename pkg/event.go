package gita

import (
	"fmt"
	"time"
)

type Event struct {
	id        int
	message   string
	timestamp time.Time
	filename  string
	line      int
}

func (e *Event) format(level Level) string {
	t := e.timestamp.Format("15:01:02")

	return fmt.Sprintf(
		"#%d %s [%s] (%s:%d) %s ",
		e.id,
		t,
		labels[level],
		e.filename,
		e.line,
		e.message,
	)
}
