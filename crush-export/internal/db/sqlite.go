package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"crush-export/internal/model"

	_ "modernc.org/sqlite"
)

// SQLiteRepository implements SessionRepository using modernc.org/sqlite.
type SQLiteRepository struct {
	db *sql.DB
}

// NewSQLiteRepository opens a read-only connection to the Crush database.
func NewSQLiteRepository(dbPath string) (*SQLiteRepository, error) {
	// Open in read-only mode
	db, err := sql.Open("sqlite", dbPath+"?mode=ro")
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &SQLiteRepository{db: db}, nil
}

// ListSessions returns all sessions ordered by creation time (newest first).
func (r *SQLiteRepository) ListSessions(ctx context.Context) ([]model.SessionSummary, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, title, message_count, created_at, updated_at 
		FROM sessions 
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("query sessions: %w", err)
	}
	defer rows.Close()

	var sessions []model.SessionSummary
	for rows.Next() {
		var s model.SessionSummary
		var createdAt, updatedAt int64
		if err := rows.Scan(&s.ID, &s.Title, &s.MessageCount, &createdAt, &updatedAt); err != nil {
			return nil, fmt.Errorf("scan session: %w", err)
		}
		s.CreatedAt = time.Unix(createdAt, 0)
		s.UpdatedAt = time.Unix(updatedAt, 0)
		sessions = append(sessions, s)
	}

	return sessions, rows.Err()
}

// GetSession returns a session with its messages by ID.
func (r *SQLiteRepository) GetSession(ctx context.Context, id string) (*model.Session, error) {
	var s model.Session
	var createdAt, updatedAt int64

	err := r.db.QueryRowContext(ctx, `
		SELECT id, title, created_at, updated_at 
		FROM sessions 
		WHERE id = ?
	`, id).Scan(&s.ID, &s.Title, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("session not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("query session: %w", err)
	}

	s.CreatedAt = time.Unix(createdAt, 0)
	s.UpdatedAt = time.Unix(updatedAt, 0)

	// Get messages
	messages, err := r.GetMessages(ctx, id)
	if err != nil {
		return nil, err
	}
	s.Messages = messages

	return &s, nil
}

// GetMessages returns all messages for a session ordered by creation time.
func (r *SQLiteRepository) GetMessages(ctx context.Context, sessionID string) ([]model.Message, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, session_id, role, parts, created_at 
		FROM messages 
		WHERE session_id = ? 
		ORDER BY created_at ASC
	`, sessionID)
	if err != nil {
		return nil, fmt.Errorf("query messages: %w", err)
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var m model.Message
		var partsJSON string
		var createdAt int64

		if err := rows.Scan(&m.ID, &m.SessionID, &m.Role, &partsJSON, &createdAt); err != nil {
			return nil, fmt.Errorf("scan message: %w", err)
		}

		m.CreatedAt = time.Unix(createdAt, 0)
		m.Content = extractTextContent(partsJSON)
		messages = append(messages, m)
	}

	return messages, rows.Err()
}

// Close releases database resources.
func (r *SQLiteRepository) Close() error {
	return r.db.Close()
}

// Part represents a message part from the JSON structure.
type part struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type textData struct {
	Text string `json:"text"`
}

type toolCallData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Input    string `json:"input"`
	Finished bool   `json:"finished"`
}

type toolResultData struct {
	ToolCallID string `json:"tool_call_id"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	IsError    bool   `json:"is_error"`
}

// extractTextContent parses the JSON parts and extracts text content.
func extractTextContent(partsJSON string) string {
	var parts []part
	if err := json.Unmarshal([]byte(partsJSON), &parts); err != nil {
		return ""
	}

	var result string
	for _, p := range parts {
		switch p.Type {
		case "text":
			var data textData
			if err := json.Unmarshal(p.Data, &data); err == nil && data.Text != "" {
				if result != "" {
					result += "\n\n"
				}
				result += data.Text
			}
		case "tool_call":
			var data toolCallData
			if err := json.Unmarshal(p.Data, &data); err == nil {
				if result != "" {
					result += "\n\n"
				}
				result += fmt.Sprintf("**Tool Call:** `%s`\n", data.Name)
				if data.Input != "" {
					// Try to pretty print JSON
					var pretty interface{}
					if err := json.Unmarshal([]byte(data.Input), &pretty); err == nil {
						if formatted, err := json.MarshalIndent(pretty, "", "  "); err == nil {
							result += fmt.Sprintf("```json\n%s\n```", string(formatted))
						} else {
							result += fmt.Sprintf("```\n%s\n```", data.Input)
						}
					} else {
						result += fmt.Sprintf("```\n%s\n```", data.Input)
					}
				}
			}
		case "tool_result":
			var data toolResultData
			if err := json.Unmarshal(p.Data, &data); err == nil {
				if result != "" {
					result += "\n\n"
				}
				status := "Result"
				if data.IsError {
					status = "Error"
				}
				result += fmt.Sprintf("**Tool %s:** `%s`\n", status, data.Name)
				if data.Content != "" {
					content := data.Content
					if len(content) > 5000 {
						content = content[:5000] + "\n... (truncated)"
					}
					result += fmt.Sprintf("```\n%s\n```", content)
				}
			}
		}
	}

	return result
}
