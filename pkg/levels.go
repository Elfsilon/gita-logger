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

var labels = []string{
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

type Level = int

func ParseLevel(name string) (int, error) {
	for i, n := range labels {
		if name == strings.ToLower(n) {
			return i, nil
		}
	}

	return InfoLevel, fmt.Errorf("no levels matching '%v' found", name)
}
