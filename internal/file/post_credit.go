package file

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/eurofurence/artshow-digital-showroom/internal/constants"
	"github.com/eurofurence/artshow-digital-showroom/internal/hash"
)

func PostCreditFor(ctx context.Context, path, title, artist, qrPath, qrURL string) string {
	hash := hash.GetMD5Hash(qrURL)

	filePath, err := MediaPath(path, "post_credits", "_"+hash+".png")
	if err != nil {
		log.Println("Could not create post_credits paths")
		return ""
	}

	if !Exist(filePath) {
		err := generatePostCredit(ctx, filePath, title, artist, qrPath)
		if err != nil {
			log.Printf("Thumbnail failed to generate: %s\n", err)
			return constants.FallbackImage
		}
	}
	return filePath
}

func generatePostCredit(ctx context.Context, outputFile, title, artist, qrPath string) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

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
				"(",
				qrPath,
				"-resize", "300x300",
				")",
				"-gravity", "southeast",
				"-geometry", "+50+50",
				"-composite",
			)
		}
	}

	// Output file
	args = append(args, outputFile)

	cmd := exec.CommandContext(timeoutCtx, "magick", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ffmpeg failed: %w\n%s", err, output)
	}
	return nil
}
