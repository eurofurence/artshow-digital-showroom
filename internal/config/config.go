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
	Thumbnail string `toml:"-"`
	Video     string `toml:"-"`
	VideoType string `toml:"-"`
	Duration  string `toml:"-"`
	ContactQR string `toml:"-"`
}
