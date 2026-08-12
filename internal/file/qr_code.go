package file

import (
	"log"

	"github.com/eurofurence/artshow-digital-showroom/internal/hash"
	"github.com/skip2/go-qrcode"
)

func QrCodeFor(videoFile, url string, border bool) (string, string) {
	if url == "" {
		return "", ""
	}

	hash := hash.GetMD5Hash(url)

	textBorder := ""
	if !border {
		textBorder = "no_border_"
	}

	filePath, fileRoute, err := MediaPaths(videoFile, "qr-codes", "_"+textBorder+hash+".png")
	if err != nil {
		log.Printf("Could not create qr code paths for %q", videoFile)
		return "", ""
	}

	if !Exist(filePath) {
		err = createQrCode(url, filePath, border)
		if err != nil {
			log.Printf("QR failed to generate: %s\n", err)
			return "", ""
		}
	}
	return filePath, fileRoute
}

func createQrCode(url, qrFile string, border bool) error {
	log.Printf("Creating qr Code %q", qrFile)

	if border {
		return qrcode.WriteFile(
			url,
			qrcode.Medium,
			256,
			qrFile,
		)
	} else {
		q, err := qrcode.New(url, qrcode.Medium)
		if err != nil {
			return err
		}

		q.DisableBorder = true

		return q.WriteFile(256, qrFile)
	}
}
