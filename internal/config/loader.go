package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Get() (*Config, error) {
	return Load("Media/config.toml")
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	s, _ := json.MarshalIndent(cfg, "", "\t")
	log.Println(string(s))

	return &cfg, nil
}
