package file

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

type Probe struct {
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

func VideoDuration(ctx context.Context, file string) (string, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		timeoutCtx,
		"ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		file,
	)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	var p Probe
	if err := json.Unmarshal(out, &p); err != nil {
		return "", err
	}

	duration, err := strconv.ParseFloat(p.Format.Duration, 64)
	if err != nil {
		return "", err
	}

	seconds := int(duration)
	return fmt.Sprintf("%02d:%02d", seconds/60, seconds%60), nil
}
