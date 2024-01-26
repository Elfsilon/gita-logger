package main

import (
	"fmt"

	gita "github.com/Elfsilon/gita-logger/pkg"
)

func main() {
	lvl, err := gita.ParseLevel("info")
	if err != nil {
		fmt.Println(err)
	}
	logger := gita.NewLogger(lvl)

	logger.Log("It's Gita's default log")
	logger.Info("It's Gita's info")
	logger.Warning("It's Gita's warning")
	logger.Error("It's Gita's error")
	logger.Fatal("It's Gita's fatal")
}
