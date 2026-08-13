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

		hash := hash.GetMD5Hash(item.File + "\x00" + item.Artist + "\x00" + item.Title)
		item.ID = hash
		s.videos[hash] = item

		videoFile := filepath.Join(s.cfg.ConfigFolder, item.File)
		item.VideoExists = file.Exist(videoFile)

		// Attributes that always exist or have fallbacks
		item.Thumbnail = file.ThumbnailFor(s.serverCtx, videoFile)

		qrPathNoBorder := file.QrCodeFor(videoFile, item.Contact, false)
		item.ContactQR = qrPathNoBorder

		qrPath := file.QrCodeFor(videoFile, item.Contact, true)
		item.PostCredit = s.cfg.MediaInterface.Address + "/" + file.PostCreditFor(
			s.serverCtx,
			videoFile,
			item.Title,
			item.Artist,
			qrPath,
			item.Contact,
		)

		// Attributes that depend on VideoExists
		if item.VideoExists {
			// Providing the local file path for Video.
			// Use the route instead when serving to a remote machine.
			absPath, _ := filepath.Abs(videoFile)
			item.Video = absPath
			// route, _ := file.MediaPath(videoFile, "", "")
			// item.Video = s.cfg.MediaInterface.Address + "/" + route

			duration, _ := file.VideoDuration(s.serverCtx, videoFile)
			item.Duration = formatDuration(duration)
			item.Preview = file.PreviewFor(s.serverCtx, videoFile, duration)
		}

	}
}
