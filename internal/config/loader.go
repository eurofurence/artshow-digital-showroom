package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"

	"github.com/eurofurence/artshow-digital-showroom/internal/arguments"
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

	preProcess(cfg)

	return cfg, nil
}

func findConfig() (string, error) {
	for _, path := range []string{"Media", "MediaExample"} {
		configPath := path + "/config.toml"

		_, err := os.Stat(configPath)
		if err == nil {
			log.Println("Using config at: " + configPath)
			return configPath, nil
		}
		log.Println("No config found at " + configPath)
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

	if arguments.Verbose() {
		s, _ := json.MarshalIndent(cfg, "", "\t")
		log.Println(string(s))
	}

	return &cfg, nil
}
