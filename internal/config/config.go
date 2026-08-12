package config

import (
	"encoding/json"
	"log"

	"github.com/eurofurence/artshow-digital-showroom/internal/arguments"
)

type Config struct {
	MediaInterface MediaInterfaceConfig `toml:"media-interface"`
	Playout        PlayoutConfig        `toml:"playout"`
	Videos         []Video              `toml:"video"`

	// Not from TOML
	ConfigFolder string `toml:"-"`
}

type MediaInterfaceConfig struct {
	Port string `toml:"port"`

	// Not from TOML
	Address string `toml:"-"`
}

type PlayoutConfig struct {
	MpvSocket string `toml:"mpv-socket"`
}

type Video struct {
	File        string `toml:"file"`
	Artist      string `toml:"artist"`
	Title       string `toml:"title"`
	Description string `toml:"info"`
	Contact     string `toml:"contact"`

	// Not from TOML
	VideoExists bool   `toml:"-"`
	ID          string `toml:"-"`
	Thumbnail   string `toml:"-"`
	Video       string `toml:"-"`
	Preview     string `toml:"-"`
	Duration    string `toml:"-"`
	ContactQR   string `toml:"-"`
	PostCredit  string `toml:"-"`
}

func (cfg Config) Print(hint string) {
	if arguments.Verbose() {
		s, _ := json.MarshalIndent(cfg, "", "\t")
		log.Println(hint, string(s))
	}
}
