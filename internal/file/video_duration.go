package file

import (
	"context"
	"encoding/json"
	"os/exec"
	"strconv"
	"time"
)

type Probe struct {
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

func VideoDuration(ctx context.Context, file string) (int, error) {
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
		return -1, err
	}

	var p Probe
	if err := json.Unmarshal(out, &p); err != nil {
		return -1, err
	}

	duration, err := strconv.ParseFloat(p.Format.Duration, 64)
	if err != nil {
		return -1, err
	}

	seconds := int(duration)
	return seconds, nil
}
