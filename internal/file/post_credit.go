package file

import (
	"fmt"
	"os"
	"os/exec"
)

func GeneratePostCredit(outputFile, title, artist, qrPath string) error {
	args := []string{
		"-size", "1920x1080",
		"xc:#000000",

		// Title
		"-font", "DejaVu-Sans",
		"-fill", "white",
		"-pointsize", "72",
		"-gravity", "north",
		"-annotate", "+0+150",
		title,

		// Artist
		"-pointsize", "42",
		"-annotate", "+0+260",
		"By " + artist,
	}

	// Add QR code only if one exists
	if qrPath != "" {
		if _, err := os.Stat(qrPath); err == nil {
			args = append(args,
				qrPath,
				"-resize", "300x300",
				"-gravity", "southeast",
				"-geometry", "+100+100",
				"-composite",
			)
		}
	}

	// Output file
	args = append(args, outputFile)

	cmd := exec.Command("magick", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ffmpeg failed: %w\n%s", err, output)
	}
	return nil
}
