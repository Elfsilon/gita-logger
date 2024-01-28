package gita

import (
	"time"
)

type Event struct {
	id         int
	message    string
	timestamp  time.Time
	filename   string
	line       int
	stackTrace string
	level      Level
}
