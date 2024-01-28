package gita

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

type Logger struct {
	Out   io.Writer
	Err   io.Writer
	file  *os.File
	level Level
	ctx   *Context
}

func NewLogger() *Logger {
	return &Logger{
		ctx:   NewContext(),
		Out:   os.Stdout,
		Err:   os.Stderr,
		level: InfoLevel,
	}
}

func (l *Logger) SetOut(out io.Writer) {
	l.Out = out
}

func (l *Logger) SetErr(err io.Writer) {
	l.Err = err
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) CreateLogFilesAt(dir string) error {
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		return err
	}

	t := time.Now().Format("2006-01-02-15:04:05.000000")
	path := fmt.Sprintf("%v/log_%v.txt", dir, t)

	l.Info(path)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, fs.ModePerm)
	if err != nil {
		return err
	}

	l.file = file
	return nil
}

func (l *Logger) Destroy() error {
	if l.file != nil {
		return l.file.Close()
	}

	return nil
}

func (l *Logger) log(message string, level Level) error {
	if level < l.level {
		return nil
	}

	event := l.ctx.NewEventFromMessage(message)
	formattedEvent := event.format(level)

	return l.write(formattedEvent)
}

func (l *Logger) write(message string) error {
	mes := message + string('\n')

	if _, err := io.WriteString(l.Out, mes); err != nil {
		return err
	}

	if l.file != nil {
		if _, err := io.WriteString(l.file, mes); err != nil {
			return err
		}
	}

	return nil
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
