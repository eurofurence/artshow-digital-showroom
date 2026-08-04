package config

type Config struct {
	MediaInterface MediaInterfaceConfig `toml:"media-interface"`
	Playout        PlayoutConfig        `toml:"playout"`
	Videos         []Video              `toml:"video"`
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
	Creator     string `toml:"creator"`
	Title       string `toml:"title"`
	Path        string `toml:"path"`
	Description string `toml:"description"`

	// Not from TOML
	Thumbnail string `toml:"-"`
}
