package config

type Config struct {
	MediaInterface MediaInterfaceConfig `toml:"media-interface"`
	Playout        PlayoutConfig        `toml:"playout"`
	Videos         []Video              `toml:"video"`

	// Not from TOML
	Folder string `toml:"-"`
}

type MediaInterfaceConfig struct {
	Port    string `toml:"port"`
	Title   string `toml:"title"`
	Columns int    `toml:"columns"`
}

type PlayoutConfig struct {
	Path string `toml:"path"`
	Port int    `toml:"port"`
}

type Video struct {
	Status      string `toml:"status"`
	File        string `toml:"file"`
	Artist      string `toml:"artist"`
	Title       string `toml:"title"`
	Description string `toml:"info"`
	Contact     string `toml:"contact"`

	// Not from TOML
	Thumbnail string `toml:"-"`
}
