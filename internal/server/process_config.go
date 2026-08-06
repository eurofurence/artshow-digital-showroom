package server

import (
	"crypto/md5"
	"encoding/hex"
	"path/filepath"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/internal/file"
)

func (s *Server) processConfig() {
	s.cfg.MediaInterface.Address = "http://localhost" + s.cfg.MediaInterface.Port
	s.videos = make(map[string]*config.Video, len(s.cfg.Videos))

	for i := range s.cfg.Videos {
		item := &s.cfg.Videos[i]

		hash := getMD5Hash(item.Title)
		item.ID = hash
		s.videos[hash] = item

		videoFile := filepath.Join(s.cfg.ConfigFolder, item.File)

		// Providing the local file path for Video.
		// Use the route instead when serving to a remote machine.
		absPath, _ := filepath.Abs(videoFile)
		item.Video = absPath
		// _, route, _ := file.MediaPaths(videoFile, "", "")
		// item.Video = s.cfg.MediaInterface.Address + "/" + route

		item.Preview = file.PreviewFor(videoFile)
		item.Duration, _ = file.VideoDuration(videoFile)

		item.Thumbnail = file.ThumbnailFor(videoFile)
		qrPath, qrRoute := file.QrCodeFor(videoFile, item.Contact)
		item.ContactQR = qrRoute

		item.PostCredit = s.cfg.MediaInterface.Address + "/" + file.PostCreditFor(
			videoFile,
			item.Title,
			item.Artist,
			qrPath,
		)
	}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
