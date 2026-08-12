package file

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/eurofurence/artshow-digital-showroom/internal/constants"
	"github.com/eurofurence/artshow-digital-showroom/internal/hash"
)

func PostCreditFor(path, title, artist, qrPath, qrURL string) string {
	hash := hash.GetMD5Hash(qrURL)
	filePath, fileRoute, err := MediaPaths(path, "post_credits", "_"+hash+".png")
	if err != nil {
		log.Println("Could not create post_credits paths")
		return ""
	}

	if !Exist(filePath) {
		err := generatePostCredit(filePath, title, artist, qrPath)
		if err != nil {
			log.Printf("Thumbnail failed to generate: %s\n", err)
			return constants.FallbackImage
		}
	}
	return fileRoute
}

func generatePostCredit(outputFile, title, artist, qrPath string) error {
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
				"-resize", "250x250",
				")",
				"-gravity", "southeast",
				"-geometry", "+50+50",
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
