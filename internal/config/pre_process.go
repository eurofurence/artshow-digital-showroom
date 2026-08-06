package config

import (
	"log"
	"mime"
	"path/filepath"

	"github.com/eurofurence/artshow-digital-showroom/internal/file"
)

func preProcess(cfg *Config) {
	cfg.MediaInterface.Address = "http://localhost" + cfg.MediaInterface.Port
	for i := range cfg.Videos {
		item := &cfg.Videos[i]

		videoFile := filepath.Join(cfg.Folder, item.File)

		item.Duration, _ = file.VideoDuration(videoFile)
		item.Thumbnail = file.ThumbnailFor(videoFile)
		item.Video = "media/" + item.File
		item.VideoType = mime.TypeByExtension(filepath.Ext(item.File))
		if item.Contact != "" {
			item.ContactQR = file.QrCodeFor(videoFile, item.Contact)
		}
	}
}
