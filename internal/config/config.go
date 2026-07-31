package config

type Config struct {
	MediaInterface MediaInterfaceConfig `toml:"media-interface"`
	PlayoutMachine PlayoutMachineConfig `toml:"playout-machine"`
	MediaItems     []MediaItem          `toml:"media-item"`
}

type MediaInterfaceConfig struct {
	Port    string `toml:"port"`
	Title   string `toml:"title"`
	Columns int    `toml:"columns"`
}

type PlayoutMachineConfig struct {
	Path string `toml:"path"`
	Port int    `toml:"port"`
}

type MediaItem struct {
	Creator     string `toml:"creator"`
	Title       string `toml:"title"`
	Path        string `toml:"path"`
	Description string `toml:"description"`
}
