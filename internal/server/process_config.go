package server

import (
	"path/filepath"

	"github.com/eurofurence/artshow-digital-showroom/internal/file"
	"github.com/eurofurence/artshow-digital-showroom/internal/hash"
)

func (s *Server) processConfig() {
	s.cfg.MediaInterface.Address = "http://localhost" + s.cfg.MediaInterface.Port

	for i := range s.cfg.Videos {
		item := &s.cfg.Videos[i]

		hash := hash.GetMD5Hash(item.Title)
		item.ID = hash
		s.videos[hash] = item

		videoFile := filepath.Join(s.cfg.ConfigFolder, item.File)

		// Providing the local file path for Video.
		// Use the route instead when serving to a remote machine.
		absPath, _ := filepath.Abs(videoFile)
		item.Video = absPath
		// _, route, _ := file.MediaPaths(videoFile, "", "")
		// item.Video = s.cfg.MediaInterface.Address + "/" + route

		item.Preview = file.PreviewFor(s.serverCtx, videoFile)
		item.Duration, _ = file.VideoDuration(s.serverCtx, videoFile)

		item.Thumbnail = file.ThumbnailFor(s.serverCtx, videoFile)
		qrPath, qrRoute := file.QrCodeFor(videoFile, item.Contact)
		item.ContactQR = qrRoute

		item.PostCredit = s.cfg.MediaInterface.Address + "/" + file.PostCreditFor(
			s.serverCtx,
			videoFile,
			item.Title,
			item.Artist,
			qrPath,
			item.Contact,
		)
	}
}
