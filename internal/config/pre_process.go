package config

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/eurofurence/artshow-digital-showroom/internal/file"
)

func preProcess(cfg *Config) {
	cfg.MediaInterface.Address = "http://localhost" + cfg.MediaInterface.Port
	for i := range cfg.Videos {
		item := &cfg.Videos[i]

		videoFile := filepath.Join(cfg.Folder, item.File)

		item.Duration, _ = videoDuration(videoFile)
		item.Thumbnail = thumbnailFor(videoFile)
		item.Video = "media/" + item.File
		item.VideoType = mime.TypeByExtension(filepath.Ext(item.File))
		if item.Contact != "" {
			item.ContactQR = file.QrCodeFor(videoFile, item.Contact)
		}
	}
}

func thumbnailFor(path string) string {
	fallback := "/static/fallback-lyca-shocked-bw.png"

	if !file.Exist(path) {
		log.Println("the video " + path + " was not found")
		return fallback
	}

	filePath, fileRoute, err := file.MediaPaths(path, "thumbnails", ".webp")
	if err != nil {
		log.Println("Could not create thumbnail paths for " + path)
		return ""
	}

	if !file.Exist(filePath) {
		err := createThumbnail(path, filePath)
		if err != nil {
			log.Printf("Thumbnail failed to generate: %s\n", err)
			return fallback
		}
	}
	return fileRoute
}

func createThumbnail(videoFile string, thumbnailFile string) error {
	log.Println("Creating thumbnail " + thumbnailFile)

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-ss", "10", // seek to x seconds
		"-i", videoFile,
		"-frames:v", "1",
		// "-vf", "scale=320:-1",
		"-c:v", "libwebp",
		"-quality", "90", // 0-100
		thumbnailFile,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ffmpeg failed: %w\n%s", err, output)
	}
	return nil
}

type Probe struct {
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

func videoDuration(file string) (string, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		file,
	)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	var p Probe
	if err := json.Unmarshal(out, &p); err != nil {
		return "", err
	}

	duration, err := strconv.ParseFloat(p.Format.Duration, 64)
	if err != nil {
		return "", err
	}

	seconds := int(duration)
	return fmt.Sprintf("%02d:%02d", seconds/60, seconds%60), nil
}
