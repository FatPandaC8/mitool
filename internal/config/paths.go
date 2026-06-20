package config

import (
	"os"
	"path/filepath"
)

func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}

func ConfigDir() string {
	return filepath.Join(HomeDir(), ".config", "mitool")
}

func ConfigFile() string {
	return filepath.Join(ConfigDir(), "config.yaml")
}

func TemplateDir() string {
	return filepath.Join(ConfigDir(), "templates")
}

func AccountsFile() string {
	return filepath.Join(ConfigDir(), "accounts.yaml")
}

func SSHConfigFile() string {
	home := HomeDir()

	return filepath.Join(home, ".ssh", "config")
}