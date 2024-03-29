package gita

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	level     Level
	Out       io.Writer
	Err       io.Writer
	ctx       *Context
	formatter *Formatter
}

func NewLogger(config *Config) *Logger {
	if config == nil {
		config = &Config{}
	}

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

	return nil
}

func (l *Logger) write_out(message string) error {
	if _, err := io.WriteString(l.Out, message+string('\n')); err != nil {
		return err
	}

	return nil
}

func (l *Logger) Log(level Level, message string) error {
	return l.log(message, level, INITIAL_CALL_DEPTH)
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

func (l *Logger) Logf(level Level, format string, a ...any) error {
	return l.log(fmt.Sprintf(format, a...), level, INITIAL_CALL_DEPTH)
}

func (l *Logger) Infof(format string, a ...any) error {
	return l.log(fmt.Sprintf(format, a...), InfoLevel, INITIAL_CALL_DEPTH)
}

func (l *Logger) Warningf(format string, a ...any) error {
	return l.log(fmt.Sprintf(format, a...), WarningLevel, INITIAL_CALL_DEPTH)
}

func (l *Logger) Errorf(format string, a ...any) error {
	return l.log(fmt.Sprintf(format, a...), ErrorLevel, INITIAL_CALL_DEPTH)
}

func (l *Logger) Fatalf(format string, a ...any) error {
	message := fmt.Sprintf(format, a...)
	if err := l.log(message, FatalLevel, INITIAL_CALL_DEPTH); err != nil {
		return err
	}
	panic(message)
}
