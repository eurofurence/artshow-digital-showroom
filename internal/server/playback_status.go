package server

import "sync"

type PlaybackStatus struct {
	mu       sync.Mutex
	idle     bool
	title    string
	position int
	duration int
}

type PlaybackStatusSnapshot struct {
	Idle     bool
	Title    string
	Position int
	Duration int
}

func (p *PlaybackStatus) SetIdle() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.idle = true
}

func (p *PlaybackStatus) IsIdle() bool {
	return p.idle
}

func (p *PlaybackStatus) StartTrack(title string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.idle = false
	p.title = title
	p.position = 0
	p.duration = 0
}

// Return true if the position now has a new value
func (p *PlaybackStatus) SetPosition(position int) bool {
	if position == p.position {
		return false
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.position = position
	return true
}

func (p *PlaybackStatus) SetDuration(duration int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.duration = duration
}

func (p *PlaybackStatus) Snapshot() PlaybackStatusSnapshot {
	p.mu.Lock()
	defer p.mu.Unlock()

	return PlaybackStatusSnapshot{
		Idle:     p.idle,
		Title:    p.title,
		Position: p.position,
		Duration: p.duration,
	}
}
