package file

import (
	"os"
	"path/filepath"
)

func Exist(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
}

// MediaPath takes a given video file and creates paths for derived media like thumbnails and qr Codes.
// The "path" is the location on disk relative to the working directory of the server.
// It also represents the location where the server serves the file, i.e., at localhost:port/<path>
func MediaPath(
	videoFile string,
	subfolder string,
	newExtension string,
) (path string, err error) {
	dir := filepath.Dir(videoFile)
	base := filepath.Base(videoFile)
	ext := filepath.Ext(base)

	var name string
	if newExtension == "" {
		name = base
	} else {
		name = base[:len(base)-len(ext)] + newExtension
	}

	path = filepath.Join(dir, subfolder, name)

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", err
	}

	return
}
