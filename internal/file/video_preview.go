package file

import (
	"fmt"
	"log"
	"os/exec"
)

func PreviewFor(path string) string {
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
		err := createPreview(path, filePath)
		if err != nil {
			log.Printf("Preview failed to generate: %s\n", err)
			return ""
		}
	}
	return fileRoute
}

func createPreview(videoFile string, previewFile string) error {
	log.Printf("Creating Preview %q", previewFile)

	cmd := exec.Command(
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
