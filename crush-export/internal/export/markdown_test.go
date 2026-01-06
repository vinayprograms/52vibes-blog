package export

import (
	"strings"
	"testing"
	"time"

	"crush-export/internal/model"
)

func TestMarkdown_BasicStructure(t *testing.T) {
	session := &model.Session{
		ID:        "test-1",
		Title:     "Test Session",
		CreatedAt: time.Date(2026, 1, 5, 10, 30, 0, 0, time.UTC),
		Messages: []model.Message{
			{ID: "m1", Role: "user", Content: "Hello", CreatedAt: time.Date(2026, 1, 5, 10, 30, 0, 0, time.UTC)},
			{ID: "m2", Role: "assistant", Content: "Hi there!", CreatedAt: time.Date(2026, 1, 5, 10, 31, 0, 0, time.UTC)},
		},
	}

	md := Markdown(session)

	// Check header
	if !strings.Contains(md, "# Session: Test Session") {
		t.Error("Missing session title")
	}

	// Check date
	if !strings.Contains(md, "**Date**: 2026-01-05 10:30:00") {
		t.Error("Missing or incorrect date")
	}

	// Check statistics table
	if !strings.Contains(md, "| Messages | 2 |") {
		t.Error("Missing message count")
	}

	// Check Lessons Learned placeholder
	if !strings.Contains(md, "## Lessons Learned") {
		t.Error("Missing Lessons Learned section")
	}
	if !strings.Contains(md, "<!-- Add your learnings here -->") {
		t.Error("Missing Lessons Learned placeholder comment")
	}

	// Check conversation log
	if !strings.Contains(md, "## Full Conversation Log") {
		t.Error("Missing Full Conversation Log section")
	}
}

func TestMarkdown_EmptySession(t *testing.T) {
	session := &model.Session{
		ID:        "empty",
		Title:     "Empty Session",
		CreatedAt: time.Now(),
	}

	md := Markdown(session)

	if !strings.Contains(md, "| Messages | 0 |") {
		t.Error("Empty session should show 0 messages")
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello-world"},
		{"Test: Special Characters!", "test-special-characters"},
		{"Multiple   Spaces", "multiple-spaces"},
		{"UPPERCASE", "uppercase"},
		{"---leading-trailing---", "leading-trailing"},
		{"", "untitled"},
		{"   ", "untitled"},
		{"123 Numbers 456", "123-numbers-456"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := slugify(tt.input)
			if result != tt.expected {
				t.Errorf("slugify(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFilename(t *testing.T) {
	session := &model.Session{
		ID:        "test-1",
		Title:     "My Test Session",
		CreatedAt: time.Date(2026, 1, 5, 14, 30, 45, 0, time.UTC),
	}

	filename := Filename(session)
	expected := "2026-01-05_14-30-45_my-test-session.md"

	if filename != expected {
		t.Errorf("Filename() = %q, want %q", filename, expected)
	}
}

func TestFilename_LongTitle(t *testing.T) {
	session := &model.Session{
		ID:        "test-1",
		Title:     "This is a very long title that should be truncated to exactly fifty characters for the filename",
		CreatedAt: time.Date(2026, 1, 5, 14, 30, 45, 0, time.UTC),
	}

	filename := Filename(session)

	// Slug should be truncated to 50 chars
	parts := strings.Split(filename, "_")
	if len(parts) < 3 {
		t.Fatalf("Filename format unexpected: %s", filename)
	}
	slug := strings.TrimSuffix(parts[2], ".md")
	if len(slug) > 50 {
		t.Errorf("Slug length %d > 50: %s", len(slug), slug)
	}
}

func TestFilenameWithSuffix(t *testing.T) {
	session := &model.Session{
		ID:        "test-1",
		Title:     "Duplicate",
		CreatedAt: time.Date(2026, 1, 5, 14, 30, 45, 0, time.UTC),
	}

	filename := FilenameWithSuffix(session, 2)
	expected := "2026-01-05_14-30-45_duplicate-2.md"

	if filename != expected {
		t.Errorf("FilenameWithSuffix() = %q, want %q", filename, expected)
	}
}

func TestTitleCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"user", "User"},
		{"assistant", "Assistant"},
		{"system", "System"},
		{"", ""},
		{"USER", "USER"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := titleCase(tt.input)
			if result != tt.expected {
				t.Errorf("titleCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSelectKeyExchanges(t *testing.T) {
	now := time.Now()
	messages := []model.Message{
		{ID: "m1", Role: "user", Content: "Start the project", CreatedAt: now},
		{ID: "m2", Role: "assistant", Content: "Sure, I'll help", CreatedAt: now.Add(time.Second)},
		{ID: "m3", Role: "user", Content: "What should we do first?", CreatedAt: now.Add(2 * time.Second)},
		{ID: "m4", Role: "assistant", Content: "Let's begin with...", CreatedAt: now.Add(3 * time.Second)},
		{ID: "m5", Role: "user", Content: "Okay", CreatedAt: now.Add(4 * time.Second)},
		{ID: "m6", Role: "assistant", Content: "Done!", CreatedAt: now.Add(5 * time.Second)},
	}

	key := selectKeyExchanges(messages)

	if len(key) == 0 {
		t.Fatal("Should return some key exchanges")
	}

	// Should include first user message
	hasFirst := false
	for _, m := range key {
		if m.ID == "m1" {
			hasFirst = true
			break
		}
	}
	if !hasFirst {
		t.Error("Key exchanges should include first user message")
	}
}

func TestSelectKeyExchanges_FewMessages(t *testing.T) {
	messages := []model.Message{
		{ID: "m1", Role: "user", Content: "Hi"},
		{ID: "m2", Role: "assistant", Content: "Hello"},
	}

	key := selectKeyExchanges(messages)

	if len(key) != 0 {
		t.Error("Should return nil for sessions with < 4 messages")
	}
}
