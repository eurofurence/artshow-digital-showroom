package config

func preProcess(cfg *Config) {
	entries := make([]Video, len(cfg.Videos))
	copy(entries, cfg.Videos)
	for i := range entries {
		entries[i].Thumbnail = thumbnailFor(entries[i].File)
	}
}

func thumbnailFor(path string) string {
	/*
		if fileExists(path) {
			generate preview image
		}
	*/

	return "/static/fallback-lyca-shocked-bw.png"
}
