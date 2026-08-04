package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func preProcess(cfg *Config) {
	for i := range cfg.Videos {
		videoFile := filepath.Join(cfg.Folder, cfg.Videos[i].File)
		cfg.Videos[i].Thumbnail = thumbnailFor(videoFile)
		cfg.Videos[i].Video = "media/" + cfg.Videos[i].File
	}
}

func thumbnailFor(path string) string {
	_, err := os.Stat(path)
	if err != nil {
		log.Println("the video " + path + " was not found")
		return "/static/fallback-lyca-shocked-bw.png"
	}

	filePath, fileRoute := thumbnailName(path)

	_, err = os.Stat(filePath)
	if err == nil {
		// thumbnail already exists
	} else {
		err = createThumbnail(path, filePath)
		if err != nil {
			log.Println("Thumbnail failed to generate: %s\n", err)
			return "/static/fallback-lyca-shocked-bw.png"
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
		"-ss", "5", // seek to 5 seconds
		"-i", videoFile,
		"-frames:v", "1",
		"-vf", "scale=320:-1",
		"-c:v", "libwebp",
		"-quality", "80", // 0-100
		thumbnailFile,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ffmpeg failed: %w\n%s", err, output)
	}
	return nil
}
