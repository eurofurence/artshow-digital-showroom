package file

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func PreviewFor(ctx context.Context, path string) string {
	if !Exist(path) {
		log.Printf("preview: the video %q was not found", path)
		return ""
	}

	filePath, fileRoute, err := MediaPaths(path, "previews", ".mp4")
	if err != nil {
		log.Printf("Could not create previews paths for %q", path)
		return ""
	}

	if !Exist(filePath) {
		err := createPreview(ctx, path, filePath)
		if err != nil {
			log.Printf("Preview failed to generate: %s\n", err)
			return ""
		}
	}
	return fileRoute
}

func createPreview(ctx context.Context, videoFile string, previewFile string) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	log.Printf("Creating Preview %q", previewFile)

	cmd := exec.CommandContext(
		timeoutCtx,
		"ffmpeg",
		"-ss", "5", // start at 5 seconds
		"-i", videoFile,
		"-t", "10", // duration 10 seconds
		"-an", // no audio
		"-c:v", "libx264",
		"-crf", "28", // constant rate factor / determines preview quality, size
		"-preset", "veryfast",
		"-pix_fmt", "yuv420p",
		"-movflags", "+faststart",
		previewFile,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ffmpeg failed: %w\n%s", err, output)
	}
	return nil
}
