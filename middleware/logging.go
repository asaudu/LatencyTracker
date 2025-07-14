package middleware

import (
	"math/rand"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json: "timestamp"`
	Level     string    `json:"level"`
	Service   string    `json:"service"`
	Message   string    `json: "message"`
}

var logs []LogEntry

func GenerateLogs() LogEntry {
	//to substitute rand.Seed since it seems deprecated
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	levels := []string{"INFO", "WARN", "ERROR"}
	services := []string{"auth", "api", "search", "stream"}
	messages := []string{
		"Request received",
		"Login failed",
		"Stream started",
		"Timeout occurred",
		"User not found",
		"Token expired",
		"Search returned results",
	}

	return LogEntry{
		Timestamp: time.Now(),
		Level:     levels[r.Intn(len(levels))],
		Service:   services[r.Intn(len(services))],
		Message:   messages[r.Intn(len(messages))],
	}
}
