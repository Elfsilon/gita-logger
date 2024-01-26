package gita

import (
	"fmt"
	"io"
	"os"
)

type Event struct{}

type Logger struct {
	Out   io.Writer
	Err   io.Writer
	level Level
}

func NewLogger(level Level) *Logger {
	return &Logger{
		Out:   os.Stdout,
		Err:   os.Stderr,
		level: level,
	}
}

func (l *Logger) log(message string, level Level) error {
	if level < l.level {
		return nil
	}
	temp := fmt.Sprintf("[%v] %v", labels[level], message)
	return l.write(temp)
}

func (l *Logger) write(message string) error {
	_, err := io.WriteString(l.Out, message+string('\n'))
	return err
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) Log(message string) error {
	return l.log(message, InfoLevel)
}

func (l *Logger) Info(message string) error {
	return l.log(message, InfoLevel)
}

func (l *Logger) Warning(message string) error {
	return l.log(message, WarningLevel)
}

func (l *Logger) Error(message string) error {
	return l.log(message, ErrorLevel)
}

func (l *Logger) Fatal(message string) error {
	return l.log(message, FatalLevel)
}
