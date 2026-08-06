package config

import (
	"path/filepath"

	"github.com/eurofurence/artshow-digital-showroom/internal/file"
)

func preProcess(cfg *Config) {
	cfg.MediaInterface.Address = "http://localhost" + cfg.MediaInterface.Port
	for i := range cfg.Videos {
		item := &cfg.Videos[i]

		videoFile := filepath.Join(cfg.ConfigFolder, item.File)

		// Providing the local file path for Video.
		// Use the route instead when serving to a remote machine.
		absPath, _ := filepath.Abs(videoFile)
		item.Video = absPath
		// _, route, _ := file.MediaPaths(videoFile, "", "")
		// item.Video = cfg.MediaInterface.Address + "/" + route

		item.Preview = file.PreviewFor(videoFile)
		item.Duration, _ = file.VideoDuration(videoFile)

		item.Thumbnail = file.ThumbnailFor(videoFile)
		item.ContactQR = file.QrCodeFor(videoFile, item.Contact)
	}
}
