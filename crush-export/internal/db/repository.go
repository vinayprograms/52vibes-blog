// Package db provides database access for crush session data.
package db

import (
	"context"

	"crush-export/internal/model"
)

// SessionRepository defines the interface for accessing session data.
type SessionRepository interface {
	// ListSessions returns all sessions ordered by creation time (newest first).
	ListSessions(ctx context.Context) ([]model.SessionSummary, error)

	// GetSession returns a session with its messages by ID.
	GetSession(ctx context.Context, id string) (*model.Session, error)

	// GetMessages returns all messages for a session ordered by creation time.
	GetMessages(ctx context.Context, sessionID string) ([]model.Message, error)

	// Close releases database resources.
	Close() error
}
