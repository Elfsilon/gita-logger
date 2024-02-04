package gita

import (
	"fmt"
	"strings"
)

const (
	InfoLevel = iota
	WarningLevel
	ErrorLevel
	FatalLevel
)

var levelNames = []string{
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

type Level = int

func ParseLevel(name string) (int, error) {
	for i, n := range levelNames {
		if strings.EqualFold(name, n) {
			return i, nil
		}
	}

	return InfoLevel, fmt.Errorf("no levels matching '%v' found", name)
}
