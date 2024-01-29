package gita

import (
	"fmt"
	"time"
)

type Event struct {
	ID         int       `json:"id,omitempty"`
	Level      Level     `json:"level"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
	Filename   string    `json:"filename,omitempty"`
	Line       int       `json:"line,omitempty"`
	StackTrace string    `json:"stackTrace,omitempty"`
}

func (e *Event) String() string {
	res := fmt.Sprintf("#%d %s [%s] (%s:%d) %s",
		e.ID,
		e.Timestamp.Format("15:01:02"),
		levelNames[e.Level],
		e.Filename,
		e.Line,
		e.Message,
	)

	if e.StackTrace != "" {
		res = fmt.Sprintf("%s\nStackTrace: %s", res, e.StackTrace)
	}

	return res
}
