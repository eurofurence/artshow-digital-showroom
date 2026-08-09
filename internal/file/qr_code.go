package file

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func QrCodeFor(videoFile string, url string) (string, string) {
	if url == "" {
		return "", ""
	}

	filePath, fileRoute, err := MediaPaths(videoFile, "qr-codes", ".png")
	if err != nil {
		log.Println("Could not create qr code paths for " + videoFile)
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

func createQrCode(url string, qrFile string) (err error) {
	log.Println("Creating qr Code " + qrFile)

	q, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return
	}

	q.DisableBorder = true

	return q.WriteFile(256, qrFile)
}
