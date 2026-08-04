package config

func preProcess(cfg *Config) {
	for i := range cfg.Videos {
		cfg.Videos[i].Thumbnail = thumbnailFor(cfg.Videos[i].File)
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
