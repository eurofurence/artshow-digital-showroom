package server

import (
	"log"
	"sync"
	"time"

	"github.com/eurofurence/artshow-digital-showroom/internal/stats"
)

type PlaybackStatus struct {
	logger    *stats.Logger
	mu        sync.Mutex
	idle      bool
	inCredits bool
	title     string
	position  int
	duration  int
}

func NewPlaybackStatus() *PlaybackStatus {
	timestamp := time.Now().Format("2006-01-02_15-04-05")

	logger, err := stats.Open("logs/log_" + timestamp + ".csv")
	if err != nil {
		log.Printf("Failed to create a logger %v", err)
	}

	return &PlaybackStatus{logger: logger}
}

type PlaybackStatusSnapshot struct {
	Idle      bool
	InCredits bool
	Title     string
	Position  int
	Duration  int
}

func (p *PlaybackStatus) SetIdle() {
	p.mu.Lock()
	defer p.mu.Unlock()

	// already logged finish if inCredits = true
	if !p.inCredits {
		p.log("aborted")
	}

	p.title = ""
	p.idle = true
	p.inCredits = false
}

func (p *PlaybackStatus) IsIdle() bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.idle
}

func (p *PlaybackStatus) StartTrack(title string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// already logged finish if inCredits = true
	if !p.inCredits {
		p.log("aborted")
	}

	p.idle = false
	p.inCredits = false
	p.title = title
	p.position = 0
	p.duration = 0

	p.log("started")
}

// Return true if the position now has a new value
func (p *PlaybackStatus) SetPosition(position int) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	if position == p.position {
		return false
	}

	p.position = position
	return true
}

func (p *PlaybackStatus) SetDuration(duration int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.duration = duration
}

func (p *PlaybackStatus) SetInCredits() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.log("finished")

	p.inCredits = true
}

func (p *PlaybackStatus) Snapshot() PlaybackStatusSnapshot {
	p.mu.Lock()
	defer p.mu.Unlock()

	return PlaybackStatusSnapshot{
		Idle:      p.idle,
		InCredits: p.inCredits,
		Title:     p.title,
		Position:  p.position,
		Duration:  p.duration,
	}
}

func (p *PlaybackStatus) Close() {
	if p.logger == nil {
		return
	}

	if err := p.logger.Close(); err != nil {
		log.Printf("csv writer failed to close %v", err)
	}
}

func (p *PlaybackStatus) log(status string) {
	if p.logger == nil || p.title == "" {
		return
	}

	if err := p.logger.Append(p.title, status); err != nil {
		log.Printf("Failed to log %v", err)
	}
}
