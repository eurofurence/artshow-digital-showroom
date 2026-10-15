package arguments

import (
	"os"
	"slices"
)

func Verbose() bool {
	return slices.Contains(os.Args, "-v")
}
