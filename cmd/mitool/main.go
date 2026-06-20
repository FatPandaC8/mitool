package main

import (
	"os"

	"github.com/FatPandaC8/mitool/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		// TODO: write usage message here
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		commands.Init()

	case "github":
		commands.Github(os.Args[2:])

	default:
		// TODO: write usage message here
		os.Exit(1)
	}
}