// Package model defines domain types for crush session exports.
package model

import (
	"time"
)

// Session represents a Crush AI conversation session.
type Session struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Messages  []Message
}

// Message represents a single message in a session.
type Message struct {
	ID        string
	SessionID string
	Role      string // "user", "assistant", "system"
	Content   string
	CreatedAt time.Time
}

// Statistics holds computed metrics for a session.
type Statistics struct {
	Duration      time.Duration
	MessageCount  int
	UserMessages  int
	AssistantMsgs int
	FilesCreated  int
	FilesModified int
	ErrorCount    int
}

// SessionSummary provides a lightweight view for listing sessions.
type SessionSummary struct {
	ID           string
	Title        string
	MessageCount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// ComputeStatistics calculates session metrics from messages.
func (s *Session) ComputeStatistics() Statistics {
	stats := Statistics{
		MessageCount: len(s.Messages),
	}

	if len(s.Messages) > 0 {
		first := s.Messages[0].CreatedAt
		last := s.Messages[len(s.Messages)-1].CreatedAt
		stats.Duration = last.Sub(first)
	}

	for _, msg := range s.Messages {
		switch msg.Role {
		case "user":
			stats.UserMessages++
		case "assistant":
			stats.AssistantMsgs++
		}
	}

	return stats
}
