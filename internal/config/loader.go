package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/eurofurence/artshow-digital-showroom/internal/file"
	"github.com/pelletier/go-toml/v2"
)

func Get() (*Config, error) {
	path, err := findConfig()
	if err != nil {
		return nil, err
	}

	cfg, err := load(path)
	if err != nil {
		return nil, err
	}

	cfg.Print("Raw loaded config")

	return cfg, nil
}

func findConfig() (string, error) {
	for _, path := range []string{"Media", "MediaExample"} {
		configPath := path + "/config.toml"

		if file.Exist(configPath) {
			log.Printf("Using config at %q", configPath)
			return configPath, nil
		}
		log.Printf("No config found at %q", configPath)
	}
	return "", errors.New("no config file found")
}

func load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	cfg.ConfigFolder = filepath.Dir(path)

	return &cfg, nil
}
