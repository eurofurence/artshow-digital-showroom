package file

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func QrCodeFor(videoFile string, url string) string {
	filePath, fileRoute, err := MediaPaths(videoFile, "qr-codes", ".webp")
	if err != nil {
		log.Println("Could not create qr code paths for " + videoFile)
		return ""
	}

	if !Exist(filePath) {
		err = createQrCode(url, filePath)
		if err != nil {
			log.Printf("QR failed to generate: %s\n", err)
			return ""
		}
	}
	return fileRoute
}

func createQrCode(url string, qrFile string) error {
	log.Println("Creating qr Code " + qrFile)

	return qrcode.WriteFile(
		url,
		qrcode.Medium,
		256,
		qrFile,
	)
}
