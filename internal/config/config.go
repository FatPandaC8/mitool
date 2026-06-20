package config

import (
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Version int `yaml:"version"`
}

func CreateDefault() error {
	cfg := Config{
		Version: 1,
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(
		ConfigFile(),
		data,
		0644,
	)
}