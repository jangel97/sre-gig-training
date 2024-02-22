package main

import (
	"os"

	"github.com/jangel97/task4/command"
)

func main() {
	cmd := command.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
