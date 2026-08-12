package file

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/eurofurence/artshow-digital-showroom/internal/constants"
)

func ThumbnailFor(path string) string {
	if !Exist(path) {
		log.Printf("thumbnail: the video %q was not found", path)
		return constants.FallbackImage
	}

	filePath, fileRoute, err := MediaPaths(path, "thumbnails", ".webp")
	if err != nil {
		log.Printf("Could not create thumbnail paths for %q", path)
		return ""
	}

	if !Exist(filePath) {
		err := createThumbnail(path, filePath)
		if err != nil {
			log.Printf("Thumbnail failed to generate: %s\n", err)
			return constants.FallbackImage
		}
	}
	return fileRoute
}

func createThumbnail(videoFile string, thumbnailFile string) error {
	log.Printf("Creating thumbnail for %q", thumbnailFile)

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
