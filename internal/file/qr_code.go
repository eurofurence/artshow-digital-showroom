package file

import (
	"log"

	"github.com/eurofurence/artshow-digital-showroom/internal/hash"
	"github.com/skip2/go-qrcode"
)

func QrCodeFor(videoFile string, url string) (string, string) {
	if url == "" {
		return "", ""
	}

	hash := hash.GetMD5Hash(url)

	filePath, fileRoute, err := MediaPaths(videoFile, "qr-codes", "_"+hash+".png")
	if err != nil {
		log.Printf("Could not create qr code paths for %q", videoFile)
		return "", ""
	}

	if !Exist(filePath) {
		err = createQrCode(url, filePath)
		if err != nil {
			log.Printf("QR failed to generate: %s\n", err)
			return "", ""
		}
	}
	return filePath, fileRoute
}

func createQrCode(url string, qrFile string) error {
	log.Printf("Creating qr Code %q", qrFile)

	return qrcode.WriteFile(
		url,
		qrcode.Medium,
		256,
		qrFile,
	)
}
