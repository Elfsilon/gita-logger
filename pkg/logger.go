package gita

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

type Logger struct {
	level     Level
	ctx       *Context
	Out       io.Writer
	Err       io.Writer
	file      *os.File
	formatter *Formatter
}

func NewLogger(config *Config) *Logger {
	if config.Out == nil {
		config.Out = os.Stdout
	}

	if config.Err == nil {
		config.Err = os.Stdout
	}

	formatter := NewDefaultFormatter()
	if config.Format != nil {
		formatter.merge(config.Format)
	}

	l := &Logger{
		ctx:       NewContext(),
		Out:       config.Out,
		Err:       config.Err,
		level:     config.Level,
		formatter: formatter,
	}

	if config.LogsDir != "" {
		l.createLogFilesAt(config.LogsDir)
	}

	return l
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

func (l *Logger) SetFormatter(formatter *Formatter) {
	l.formatter = formatter
}

func (l *Logger) log(message string, level Level, depth int) error {
	if level < l.level {
		return nil
	}

	event := l.ctx.NewEventFromMessage(message, level, l.formatter.DisplayStackTrace, depth+1)
	formatted := l.formatter.Format(event)

	if err := l.write_out(formatted); err != nil {
		return err
	}

	if l.file != nil {
		if err := l.write_file(event.String()); err != nil {
			return err
		}
	}

	return nil
}

func (l *Logger) write_out(message string) error {
	if _, err := io.WriteString(l.Out, message+string('\n')); err != nil {
		return err
	}

	return nil
}

func (l *Logger) write_file(message string) error {
	if _, err := io.WriteString(l.file, message+string('\n')); err != nil {
		return err
	}

	return nil
}

func (l *Logger) createLogFilesAt(dir string) error {
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		return err
	}

	t := time.Now().Format("2006-01-02-15:04:05.000")
	path := fmt.Sprintf("%v/log_%v.log", dir, t)

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

func (l *Logger) Log(message string) error {
	return l.log(message, InfoLevel, INITIAL_CALL_DEPTH)
}

func (l *Logger) Info(message string) error {
	return l.log(message, InfoLevel, INITIAL_CALL_DEPTH)
}

func (l *Logger) Warning(message string) error {
	return l.log(message, WarningLevel, INITIAL_CALL_DEPTH)
}

func (l *Logger) Error(message string) error {
	return l.log(message, ErrorLevel, INITIAL_CALL_DEPTH)
}

func (l *Logger) Fatal(message string) error {
	if err := l.log(message, FatalLevel, INITIAL_CALL_DEPTH); err != nil {
		return err
	}
	panic(message)
}
