package config

import (
	"os"
	"log"
	"encoding/json"

	"github.com/pelletier/go-toml/v2"
)

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
