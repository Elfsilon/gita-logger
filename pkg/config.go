package gita

import "io"

const INITIAL_CALL_DEPTH = 2

type Config struct {
	LogsDir string
	Level   Level
	Out     io.Writer
	Err     io.Writer
	Format  *Formatter
}
