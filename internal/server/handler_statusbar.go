package server

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Status struct {
	Idle     bool
	Title    string
	Position float64
	Duration float64
}

func (s *Server) StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	// One channel per connected client.
	ch := make(chan string, 8)

	s.clientsMu.Lock()
	s.clients[ch] = struct{}{}
	s.clientsMu.Unlock()

	defer func() {
		s.clientsMu.Lock()
		delete(s.clients, ch)
		s.clientsMu.Unlock()
		close(ch)
	}()

	for {
		select {
		case <-r.Context().Done():
			log.Println("SSE client disconnected")
			return

		case html := <-ch:
			log.Println(html)
			if err := writeSSE(w, html); err != nil {
				log.Printf("SSE write failed: %v", err)
				return
			}
			flusher.Flush()
		}
	}
}

func (s *Server) PublishStatus(status Status) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, "statusbar", status); err != nil {
		log.Printf("Error rendering statusbar in publish %v", err)
	}
	html := buf.String()

	for ch := range s.clients {
		select {
		case ch <- html: // delivered
		default: // Client is slow. Drop this update.
		}
	}
}

func writeSSE(w io.Writer, data string) error {
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		if _, err := fmt.Fprintf(w, "data: %s\n", scanner.Text()); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	_, err := fmt.Fprint(w, "\n")
	return err
}

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
