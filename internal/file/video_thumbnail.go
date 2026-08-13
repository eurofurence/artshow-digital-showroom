package file

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/eurofurence/artshow-digital-showroom/internal/constants"
)

func ThumbnailFor(ctx context.Context, path string) string {
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
		err := createThumbnail(ctx, path, filePath)
		if err != nil {
			log.Printf("Thumbnail failed to generate: %s\n", err)
			return constants.FallbackImage
		}
	}
	return fileRoute
}

func createThumbnail(ctx context.Context, videoFile string, thumbnailFile string) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	log.Printf("Creating thumbnail for %q", thumbnailFile)

	cmd := exec.CommandContext(
		timeoutCtx,
		"ffmpeg",
		"-y",
		"-ss", "10", // seek to x seconds
		"-i", videoFile,
		"-frames:v", "1",
		"-vf", "scale=480:-1", // scale to 480px width and same aspect ratio
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
