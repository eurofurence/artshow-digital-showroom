package file

import (
	"fmt"
	"log"
	"os/exec"
)

func ThumbnailFor(path string) string {
	fallback := "/static/fallback-lyca-shocked-bw.png"

	if !Exist(path) {
		log.Println("the video " + path + " was not found")
		return fallback
	}

	filePath, fileRoute, err := MediaPaths(path, "thumbnails", ".webp")
	if err != nil {
		log.Println("Could not create thumbnail paths for " + path)
		return ""
	}

	if !Exist(filePath) {
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
