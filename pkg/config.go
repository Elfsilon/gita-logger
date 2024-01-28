package gita

import "io"

const INITIAL_CALL_DEPTH = 2

type Config struct {
	DisplayID          bool
	DisplayTime        bool
	DisplayFileAndLine bool
	DisplayStackTrace  bool
	LogsDir            string
	Level              Level
	Out                io.Writer
	Err                io.Writer
}
