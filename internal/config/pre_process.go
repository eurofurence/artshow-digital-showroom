package config

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
	}
}

func thumbnailFor(path string) string {
	fallback := "/static/fallback-lyca-shocked-bw.png"

	_, err := os.Stat(path)
	if err != nil {
		log.Println("the video " + path + " was not found")
		return fallback
	}

	filePath, fileRoute := thumbnailName(path)

	_, err = os.Stat(filePath)
	if err == nil {
		// thumbnail already exists
	} else {
		err = createThumbnail(path, filePath)
		if err != nil {
			log.Printf("Thumbnail failed to generate: %s\n", err)
			return fallback
		}
		// thumbnail was generated
	}
	return fileRoute
}

func thumbnailName(videoFile string) (string, string) {
	dir := filepath.Dir(videoFile)
	base := filepath.Base(videoFile)
	ext := filepath.Ext(base)

	name := base[:len(base)-len(ext)]

	return filepath.Join(dir, "thumbnails", name+".webp"),
		filepath.Join("media", "thumbnails", name+".webp")
}

func createThumbnail(videoFile, thumbnailFile string) error {
	if err := os.MkdirAll(filepath.Dir(thumbnailFile), 0o755); err != nil {
		return err
	}

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
