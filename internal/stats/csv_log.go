package stats

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger struct {
	mu   sync.Mutex
	file *os.File
	csv  *csv.Writer
}

func Open(path string) (*Logger, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}

	_, err := os.Stat(path)

	newFile := os.IsNotExist(err)
	if err != nil && !newFile {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}

	logger := &Logger{
		file: file,
		csv:  csv.NewWriter(file),
	}

	if newFile {
		if err := logger.csv.Write([]string{
			"timestamp",
			"video",
			"status",
		}); err != nil {
			if closeErr := file.Close(); closeErr != nil {
				return nil, fmt.Errorf("write header: %w; close file: %v", err, closeErr)
			}
			return nil, err
		}

		logger.csv.Flush()

		if err := logger.csv.Error(); err != nil {
			if closeErr := file.Close(); closeErr != nil {
				return nil, fmt.Errorf("flush header: %w; close file: %v", err, closeErr)
			}
			return nil, err
		}
	}

	return logger, nil
}

func (l *Logger) Append(video, status string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format(time.RFC3339)

	if err := l.csv.Write([]string{
		timestamp,
		video,
		status,
	}); err != nil {
		return err
	}

	l.csv.Flush()
	return l.csv.Error()
}

func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.csv.Flush()

	if err := l.csv.Error(); err != nil {
		if closeErr := l.file.Close(); closeErr != nil {
			return fmt.Errorf("flush CSV: %w; close file: %v", err, closeErr)
		}
		return err
	}

	return l.file.Close()
}
