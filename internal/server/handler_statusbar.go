package server

import (
	"fmt"
	"time"
)

func formatDuration(seconds float64) string {
	if seconds <= 0 {
		return "--:--"
	}
	d := time.Duration(seconds * float64(time.Second))

	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("%d:%02d", m, s)
}
