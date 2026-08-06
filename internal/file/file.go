package file

import (
	"os"
)

func Exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
