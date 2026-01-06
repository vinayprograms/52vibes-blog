package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

func TestNewSQLiteRepository_FileNotFound(t *testing.T) {
	_, err := NewSQLiteRepository("/nonexistent/path/crush.db")
	if err == nil {
		t.Fatal("expected error for nonexistent database")
	}
}

func TestSQLiteRepository_Integration(t *testing.T) {
	// Create test database
	dbPath := filepath.Join(t.TempDir(), "test.db")
	setupTestDB(t, dbPath)

	repo, err := NewSQLiteRepository(dbPath)
	if err != nil {
		t.Fatalf("NewSQLiteRepository: %v", err)
	}
	defer repo.Close()

	ctx := context.Background()

	// Test ListSessions
	sessions, err := repo.ListSessions(ctx)
	if err != nil {
		t.Fatalf("ListSessions: %v", err)
	}
	if len(sessions) != 2 {
		t.Errorf("ListSessions returned %d sessions, want 2", len(sessions))
	}

	// Verify order (newest first)
	if len(sessions) >= 2 && sessions[0].ID != "session-2" {
		t.Errorf("Sessions not ordered by created_at DESC: first=%s", sessions[0].ID)
	}

	// Test GetSession
	session, err := repo.GetSession(ctx, "session-1")
	if err != nil {
		t.Fatalf("GetSession: %v", err)
	}
	if session.Title != "Test Session 1" {
		t.Errorf("Session title = %q, want %q", session.Title, "Test Session 1")
	}
	if len(session.Messages) != 2 {
		t.Errorf("Session has %d messages, want 2", len(session.Messages))
	}

	// Test GetSession not found
	_, err = repo.GetSession(ctx, "nonexistent")
	if err == nil {
		t.Fatal("GetSession should error for nonexistent session")
	}

	// Test GetMessages
	messages, err := repo.GetMessages(ctx, "session-1")
	if err != nil {
		t.Fatalf("GetMessages: %v", err)
	}
	if len(messages) != 2 {
		t.Errorf("GetMessages returned %d messages, want 2", len(messages))
	}
	if messages[0].Role != "user" {
		t.Errorf("First message role = %q, want %q", messages[0].Role, "user")
	}
}

func TestExtractTextContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains string
	}{
		{
			name:     "simple text",
			input:    `[{"type":"text","data":{"text":"Hello world"}}]`,
			contains: "Hello world",
		},
		{
			name:     "tool call",
			input:    `[{"type":"tool_call","data":{"id":"1","name":"bash","input":"{\"cmd\":\"ls\"}"}}]`,
			contains: "Tool Call",
		},
		{
			name:     "tool result",
			input:    `[{"type":"tool_result","data":{"tool_call_id":"1","name":"bash","content":"file.txt"}}]`,
			contains: "Tool Result",
		},
		{
			name:     "invalid json",
			input:    `not json`,
			contains: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractTextContent(tt.input)
			if tt.contains != "" && !contains(result, tt.contains) {
				t.Errorf("extractTextContent(%q) = %q, want to contain %q", tt.input, result, tt.contains)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && searchSubstring(s, substr)))
}

func searchSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func setupTestDB(t *testing.T, dbPath string) {
	t.Helper()

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		t.Fatalf("open test db: %v", err)
	}
	defer db.Close()

	// Create schema
	_, err = db.Exec(`
		CREATE TABLE sessions (
			id TEXT PRIMARY KEY,
			parent_session_id TEXT,
			title TEXT NOT NULL,
			message_count INTEGER DEFAULT 0,
			created_at INTEGER NOT NULL,
			updated_at INTEGER NOT NULL
		);
		CREATE TABLE messages (
			id TEXT PRIMARY KEY,
			session_id TEXT NOT NULL,
			role TEXT NOT NULL,
			parts TEXT NOT NULL,
			model TEXT,
			created_at INTEGER NOT NULL,
			FOREIGN KEY (session_id) REFERENCES sessions(id)
		);
	`)
	if err != nil {
		t.Fatalf("create schema: %v", err)
	}

	now := time.Now().Unix()

	// Insert test sessions
	_, err = db.Exec(`
		INSERT INTO sessions (id, title, message_count, created_at, updated_at) VALUES
		('session-1', 'Test Session 1', 2, ?, ?),
		('session-2', 'Test Session 2', 1, ?, ?)
	`, now-3600, now-3600, now, now)
	if err != nil {
		t.Fatalf("insert sessions: %v", err)
	}

	// Insert test messages
	_, err = db.Exec(`
		INSERT INTO messages (id, session_id, role, parts, created_at) VALUES
		('msg-1', 'session-1', 'user', '[{"type":"text","data":{"text":"Hello"}}]', ?),
		('msg-2', 'session-1', 'assistant', '[{"type":"text","data":{"text":"Hi there!"}}]', ?),
		('msg-3', 'session-2', 'user', '[{"type":"text","data":{"text":"Test"}}]', ?)
	`, now-3600, now-3500, now)
	if err != nil {
		t.Fatalf("insert messages: %v", err)
	}

	// Sync to disk
	if err := db.Close(); err != nil {
		t.Fatalf("close db: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(dbPath); err != nil {
		t.Fatalf("db file not created: %v", err)
	}
}
