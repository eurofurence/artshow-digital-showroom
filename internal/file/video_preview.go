package file

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func PreviewFor(ctx context.Context, path string, playtime int) string {
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
		err := createPreview(ctx, path, filePath, playtime)
		if err != nil {
			log.Printf("Preview failed to generate: %s\n", err)
			return ""
		}
	}
	return fileRoute
}

func createPreview(ctx context.Context, videoFile string, previewFile string, playtime int) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	log.Printf("Creating Preview %q", previewFile)

	var seek, duration string
	if playtime >= 15 {
		seek = "5"
		duration = "10"
	} else {
		seek = "0"
		duration = strconv.Itoa(playtime)
	}

	cmd := exec.CommandContext(
		timeoutCtx,
		"ffmpeg",
		"-ss", seek,
		"-i", videoFile,
		"-t", duration,
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
