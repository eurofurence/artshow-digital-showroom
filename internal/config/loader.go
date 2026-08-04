package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Get() (*Config, error) {
	_, _ = findConfig()
	return load("Media/config.toml")
}

func findConfig() (string, error) {
	_, err := os.Stat("Media/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	return "", nil

}

func load(path string) (*Config, error) {
	_, err := os.Stat("Media/config.toml")
	if err != nil {
		log.Fatal(err)
	}

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
