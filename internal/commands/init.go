package commands

import (
	"fmt"
	"os"

	"github.com/FatPandaC8/mitool/internal/config"
)

func Init() {
	fmt.Println("[init] mitool")

	err := os.MkdirAll(
		config.ConfigDir(),
		0755,
	)

	if err != nil {
		fmt.Println("[init] failed:", err)
		os.Exit(1)
	}

	if _, err := os.Stat(config.ConfigFile()); os.IsNotExist(err) {
		if err := config.CreateDefault(); err != nil {
			fmt.Println("[init] failed:", err)
			os.Exit(1)
		}

		fmt.Println("[init] created config")
	} else {
		fmt.Println("[init] config already exist")
	}

	fmt.Println("[init] ready at:", config.ConfigDir())
}